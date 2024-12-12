package lock_free

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type StackInterface interface {
	Push(any)
	Pop() any
}

func NewStackBySlice() *StackBySlice {
	return &StackBySlice{lock: &sync.Mutex{}}
}

type StackBySlice struct {
	data []any
	lock *sync.Mutex
}

func (s *StackBySlice) Push(item any) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, item)
}

func (s *StackBySlice) Pop() any {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

type StackByLockFree struct {
	// stack top
	top unsafe.Pointer
	// stack lens
	len uint64
}

type directItem struct {
	next unsafe.Pointer
	v    interface{}
}

func NewStackByLockFree() *StackByLockFree {
	return &StackByLockFree{}
}
func (s *StackByLockFree) Push(v any) {
	item := directItem{v: v}
	var top unsafe.Pointer

	for {
		top = atomic.LoadPointer(&s.top)
		item.next = top
		if atomic.CompareAndSwapPointer(&s.top, top, unsafe.Pointer(&item)) {
			atomic.AddUint64(&s.len, 1)
			return
		}
	}
}
func (s *StackByLockFree) Pop() any {
	var top, next unsafe.Pointer
	var item *directItem
	for {
		top = atomic.LoadPointer(&s.top)
		if top == nil {
			return nil
		}
		item = (*directItem)(top)
		next = atomic.LoadPointer(&item.next)
		if atomic.CompareAndSwapPointer(&s.top, top, next) {
			atomic.AddUint64(&s.len, ^uint64(0))
			return item.v
		}
	}
}
