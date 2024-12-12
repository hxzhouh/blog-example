package lock_free

import (
	"testing"
)

func TestNewStackByLockFree(t *testing.T) {
	stack := NewStackByLockFree()
	if stack == nil {
		t.Errorf("NewStackByLockFree() failed")
	}
	push(t, stack)
}

func TestNewStackBySlice(t *testing.T) {
	stack := NewStackBySlice()
	if stack == nil {
		t.Errorf("NewStackBySlice() failed")
	}
	push(t, stack)
}

func BenchmarkConcurrentPushLockFree(b *testing.B) {
	stack := NewStackByLockFree()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkConcurrentPushSlice(b *testing.B) {
	stack := NewStackBySlice()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func push(t *testing.T, s StackInterface) {
	item := []int{1, 2, 3, 4}
	for _, i := range item {
		s.Push(i)
	}
	for i := 0; i < len(item); i++ {
		if s.Pop() != item[len(item)-i-1] {
			t.Errorf("pop failed")
		}
	}
}
