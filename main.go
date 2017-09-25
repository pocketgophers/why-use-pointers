package main

import (
	"fmt"
	"testing"
)

type BigStruct struct {
	value int
	stuff [1000000]int
}

func main() {
	// benchmark multiple times for benchstat
	for i := 0; i < 5; i++ {
		result := testing.Benchmark(func(b *testing.B) {
			big := create()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				call(big)
			}
		})

		fmt.Printf("BenchmarkCall\t%s\t%s\n",
			result, result.MemString())
	}
}
