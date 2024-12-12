package main

import (
	"fmt"
	"time"
)

func Example() {
	a := 1
	c := make(chan error)
	defer close(c)
	go func() {
		var err error
		c <- err
		return
	}()
	// Example exits here, causing a goroutine leak.
	if a > 0 {
		return
	}
	err := <-c
	fmt.Println(err)
}

func main() {
	Example()
	time.Sleep(1 * time.Second)
}
