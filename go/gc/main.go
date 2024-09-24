package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	var stringBytes []byte
	var stringOffsets []int

	for i := 0; i < 1e8; i++ {
		val := strconv.Itoa(i)

		stringBytes = append(stringBytes, val...)
		stringOffsets = append(stringOffsets, len(stringBytes))
	}

	runtime.GC()
	start := time.Now()
	runtime.GC()
	fmt.Printf("GC took %s\n", time.Since(start))

	sStart := 0
	for i := 0; i < 10; i++ {
		sEnd := stringOffsets[i]
		bytes := stringBytes[sStart:sEnd]
		stringVal := *(*string)(unsafe.Pointer(&bytes))
		fmt.Println(stringVal)

		sStart = sEnd
	}
}
