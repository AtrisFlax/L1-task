package benchmark

import (
	squaresSum "task3"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squaresSum.CalculatStreams([]int{2, 4, 6, 8, 10})
	}
}
