package main

import "testing"

func BenchmarkFreq(b *testing.B) {
	//b.ReportAllocs()

	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, "index.xml")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freq(files)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	//b.ReportAllocs()

	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, "index.xml")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrent(files)
	}
}

func BenchmarkSyncPool(b *testing.B) {
	//b.ReportAllocs()

	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, "index.xml")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		syncPool(files)
	}
}
