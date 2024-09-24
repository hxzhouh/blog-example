package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	time.Sleep(1 * time.Second)
	for {
		go func() {
			time.Sleep(1 * time.Hour)
		}()
		time.Sleep(50 * time.Millisecond)
	}
}
