package benchmark

import (
	squares "task2"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squares.CalculateStreams([]int{2, 4, 6, 8, 10})
	}
}
