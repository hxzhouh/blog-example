package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()

	time.Sleep(time.Second * 1000)
}
