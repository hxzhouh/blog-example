package limit_alg

import (
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mutex          sync.Mutex
	counters       []int
	limit          int
	windowStart    time.Time
	windowDuration time.Duration
	interval       time.Duration
}

func NewSlidingWindowLimiter(limit int, windowDuration time.Duration, interval time.Duration) *SlidingWindowLimiter {
	buckets := int(windowDuration / interval)
	l := &SlidingWindowLimiter{
		counters:       make([]int, buckets),
		limit:          limit,
		windowStart:    time.Now(),
		windowDuration: windowDuration,
		interval:       interval,
	}
	go l.reset()
	return l
}

func (s *SlidingWindowLimiter) reset() {
	ticker := time.NewTicker(s.interval)
	for {
		select {
		case <-ticker.C:
			s.slideWindow()
		}
	}
}

// Allow 方法用于判断当前请求是否被允许，并实现滑动窗口的逻辑。
func (s *SlidingWindowLimiter) Allow() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	index := int((now.UnixNano()-s.windowStart.UnixNano())/s.interval.Nanoseconds()) % len(s.counters)
	if s.counters[index] < s.limit {
		s.counters[index]++
		return true
	}
	return false
}

func (s *SlidingWindowLimiter) slideWindow() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 滑动窗口，忽略最旧的时间段
	copy(s.counters, s.counters[1:])
	// 重置最后一个时间段的计数器
	s.counters[len(s.counters)-1] = 0
	// 更新窗口开始时间
	s.windowStart = time.Now()
}
