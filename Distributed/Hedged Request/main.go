package main

import (
	"context"
	"fmt"
	"log"
	rand2 "math/rand/v2"
	"net/http"
	"time"
)

var ( // there URLs represent repicas of the same service that we want to hedge
	client = http.DefaultClient // client is safe for concurrent use so define only once globally
)

func main() {
	go func() {
		http.HandleFunc("/", Backends)
		http.ListenAndServe(":8090", nil)
	}()
	fmt.Println("got response:", hedgedRequest())
}

func Backends(w http.ResponseWriter, r *http.Request) {
	// P99 latency of 1s
	s := time.Duration(rand2.IntN(1000)) * time.Millisecond
	time.Sleep(s)
	w.WriteHeader(http.StatusOK)
}

// hedgedRequest takes a slice of urls, invokes all of them parallely and returns the first response body received. It cancels the other inflight requests to save resources
func hedgedRequest() string {
	ch := make(chan string) // chan used to abort other requests
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(ctx *context.Context, ch chan string, i int) {
			log.Println("in goroutine: ", i)
			if request(ctx, "http://localhost:8090", i) {
				ch <- fmt.Sprintf("finsh [from %v]", i)
				log.Println("completed goroutine: ", i)
			}
		}(&ctx, ch, i)
	}

	select {
	case s := <-ch:
		cancel()
		log.Println("cancelled all inflight requests")
		return s
	case <-time.After(5 * time.Second):
		cancel()
		return "all requests timeout after 5 secs"
	}
}

func request(ctx *context.Context, url string, goroutine int) bool {
	ch := make(chan struct{})
	defer close(ch)
	go func() {
		http.NewRequestWithContext(*ctx, "GET", url, nil)
		ch <- struct{}{}
	}()
	select {
	case <-(*ctx).Done():
		log.Println("request cancelled for goroutine: ", goroutine)
		return false
	case <-ch:
		log.Println("request completed for goroutine: ", goroutine)
		return true
	}
}
