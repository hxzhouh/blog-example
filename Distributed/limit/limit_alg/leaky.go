package limit_alg

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	queue chan struct{} // 请求队列
}

// NewLeakyBucketLimit 创建一个新的漏桶实例
func NewLeakyBucketLimit(limit int) *LeakyBucket {
	lb := &LeakyBucket{
		queue: make(chan struct{}, limit),
	}
	go lb.process()
	return lb
}

func (lb *LeakyBucket) Allow() bool {
	// 如果通道可以发送，请求被接受
	return lb.push()
}

func (lb *LeakyBucket) push() bool {
	// 如果通道可以发送，请求被接受
	select {
	case lb.queue <- struct{}{}:
		return true
	default:
		return false
	}
}

func (lb *LeakyBucket) process() {
	for range lb.queue { // 使用 range 来持续接收队列中的请求
		fmt.Println("Request processed at", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(100 * time.Millisecond) // 模拟请求处理时间
	}
}
