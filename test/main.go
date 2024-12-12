package main

import (
	"fmt"
	"time"
)

//func demo() {
//	// GetLock
//	for {
//		if getRedisLock(ctx, key, duration) {
//			// do something
//			break
//		} else {
//			time.Sleep()
//
//		}
//	}
//	defer func() {
//		//free lock
//	}()
//	// 定时每5s 执行一次 renewLock
//	time.AfterFunc(5*time.Second, func() {})
//
//}

func main() {
	for {
		time.AfterFunc(5*time.Second, func() {
			fmt.Println("now time:", time.Now())
		})
	}
}
