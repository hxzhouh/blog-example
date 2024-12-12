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

// Allow
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
	// Slide the window to ignore the oldest time period
	copy(s.counters, s.counters[1:])
	// reset the last counter
	s.counters[len(s.counters)-1] = 0
	// window start  time
	s.windowStart = time.Now()
}
