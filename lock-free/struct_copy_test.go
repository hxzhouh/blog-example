package lock_free

import (
	"sync"
	"testing"
)

func BenchmarkSingleBuffer_print(b *testing.B) {

	for i := 0; i < b.N; i++ {
		run_buff_bench(singleBuff)
	}
}

func BenchmarkBufferWrapper_print(b *testing.B) {
	bufferWrapper := NewBufferWrapper()
	for i := 0; i < b.N; i++ {
		run_buff_bench(bufferWrapper)
	}
}

func run_buff_bench(buff printId) {
	wg := sync.WaitGroup{}
	f := func(id int) {
		defer wg.Done()
		for j := 0; j < 10000; j++ {
			buff.print(id)
		}
	}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go f(j)
	}
	wg.Wait()
}

func BenchmarkStructCopy_print(b *testing.B) {

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		f := func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				singleBuff.print(id)
			}
		}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go f(i)
		}
	}
}
