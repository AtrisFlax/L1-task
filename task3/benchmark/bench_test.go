package benchmark

import (
	squaresSum "task3"
	"testing"
)

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squaresSum.Calculate([]int{2, 4, 6, 8, 10})
	}
}
