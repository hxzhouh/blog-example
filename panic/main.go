package main

import "fmt"

func main() {
	a := make([]int, 0)
	for i := 0; i < 5; i++ {
		a = append(a, i)
	}
	a = a[:0]
	fmt.Println(a)
}
