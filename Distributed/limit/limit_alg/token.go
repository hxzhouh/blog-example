package limit_alg

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu         sync.Mutex
	bucketSize int
	capacity   chan struct{}
	interval   time.Duration
	refillRate int
}

func NewTokenBucket(capacity int, refillRate int, interval time.Duration) *TokenBucket {
	tl := &TokenBucket{
		capacity:   make(chan struct{}, capacity),
		bucketSize: capacity,
		interval:   interval,
		refillRate: refillRate,
	}
	go tl.refill()
	return tl
}

// Allow checks if the token bucket allows a request.
func (t *TokenBucket) Allow() bool {
	select {
	case <-t.capacity:
		return true
	default:
		return false
	}
}

func (t *TokenBucket) refill() {
	ticker := time.NewTicker(t.interval)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < t.refillRate; i++ {
				// If the channel is full, It is not a rigorous realization.
				if len(t.capacity) >= cap(t.capacity) {
					continue
				}
				t.capacity <- struct{}{}
			}
		}
	}
}
