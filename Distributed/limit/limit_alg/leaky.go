package limit_alg

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	queue chan struct{} // Use a channel to store requests
}

// NewLeakyBucketLimit creates a new LeakyBucket instance.
func NewLeakyBucketLimit(limit int) *LeakyBucket {
	lb := &LeakyBucket{
		queue: make(chan struct{}, limit),
	}
	go lb.process()
	return lb
}

func (lb *LeakyBucket) Allow() bool {
	return lb.push()
}

func (lb *LeakyBucket) push() bool {
	select {
	case lb.queue <- struct{}{}:
		return true
	default:
		return false
	}
}

func (lb *LeakyBucket) process() {
	for range lb.queue { // Use range to continuously receive requests from the queue
		fmt.Println("Request processed at", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(100 * time.Millisecond) // Simulate request processing time
	}
}
