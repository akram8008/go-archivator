package main

import (
	"testing"
)

func Benchmark_sequencedArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		sequencedArchivator([]string{
			"test1.txt",
			"test2.txt",
			"test3.txt",
			"test1.txt",
			"test2.txt",
			"test3.txt",
			"test1.txt",
			"test2.txt",
			"test3.txt",
			"test1.txt",
			"test2.txt",
			"test3.txt",
		})
	}
}

func Benchmark_concurrentArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		concurrentArchivator([]string{
			"test1.txt",
			"test1.txt",
			"test2.txt",
			"test3.txt",
			"test2.txt",
			"test1.txt",
			"test2.txt",
			"test3.txt",
			"test1.txt",
			"test3.txt",
			"test2.txt",
			"test3.txt",
			"test1.txt",
			"test3.txt",
			"test2.txt",
		})
	}
}