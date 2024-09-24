package limit_alg

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	mu       sync.Mutex
	count    int
	limit    int
	duration time.Duration
}

func NewFixedWindowCounter(limit int, duration time.Duration) *FixedWindowCounter {
	l := &FixedWindowCounter{
		limit:    limit,
		duration: duration, // Setting the duration of the time window。
	}
	go l.resetLimit()
	return l
}

func (f *FixedWindowCounter) resetLimit() {
	ticker := time.NewTicker(f.duration)
	for {
		select {
		case <-ticker.C:
			f.mu.Lock()
			f.count = 0
			f.mu.Unlock()
		}
	}
}

func (f *FixedWindowCounter) Allow() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	// allowed request
	if f.count < f.limit {
		f.count++
		return true
	}
	//	reject request
	return false
}
