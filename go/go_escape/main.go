package main

import (
	"fmt"
	"unsafe"
)

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func main() {
	v := "Hello,World"
	v2 := "Hello,World1"
	fmt.Printf("addr of v in bar = %p \n", (*int)(noescape(unsafe.Pointer(&v))))
	fmt.Printf("addr of v2 in bar = %p\n", &v2)
}
