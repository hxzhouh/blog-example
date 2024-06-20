package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	A int
	B string
}

func main() {
	fmt.Println(unsafe.Sizeof(Test{}))
	fmt.Println(unsafe.Sizeof(struct{}{}))
}
