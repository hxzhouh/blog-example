package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call the function we're benchmarking
		fib(20)
	}
}

func BenchmarkFib28(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call the function we're benchmarking
		fib(28)
	}
}

var nums []int

// bubbleSort
func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// quickSort
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	head, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}

func softNums() {
	//bubbleSort(nums)
	//quickSort(nums)
}

func initNums(count int) {
	rand.Seed(time.Now().UnixNano())
	nums = make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Intn(count)
	}
}
func BenchmarkSoftNums(b *testing.B) {
	initNums(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		softNums()
	}
}
