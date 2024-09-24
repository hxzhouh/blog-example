package main

import (
	"blog-example/Distributed/limit/limit_alg"
	"fmt"
	"time"
)

func main() {
	//limiter := limit_alg.NewFixedWindowCounter(10, time.Second)
	//limiter := limit_alg.NewSlidingWindowLimiter(1, time.Second, 100*time.Millisecond)
	limiter := limit_alg.NewTokenBucket(10, 1, 100*time.Millisecond)
	time.Sleep(900 * time.Millisecond)
	//for i := 0; i < 10; i++ {
	//	if limiter.Allow() {
	//		fmt.Println("Request", i+1, "allowed")
	//	} else {
	//		fmt.Println("Request", i+1, "rejected")
	//	}
	//}
	//time.Sleep(200 * time.Millisecond)
	for i := 0; i < 1000; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rejected")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
