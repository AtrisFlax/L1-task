package benchmark

import (
	squares "task2"
	"testing"
)

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squares.Calculate([]int{2, 4, 6, 8, 10})
	}
}
