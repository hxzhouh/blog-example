package main

import "fmt"

func main() {
	a := make([]int, 0)
	b := append(a, 1)
	c := append(a, 2)
	fmt.Println(a, b, c)
}
