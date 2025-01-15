package main

import (
	"fmt"
	"weak"
)

func main() {
	age := int64(32)
	age2 := int64(32)
	_ = weak.Make(&age)
	if age != age2 {
		fmt.Printf("%d %d\n", age, age2)

	}
}
