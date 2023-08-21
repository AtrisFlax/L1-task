package squaresSum

import (
	"github.com/jucardi/go-streams/streams"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных
вычислений.
*/

func CalculatStreams(inputNums []int) int {
	sum := 0
	streams.
		From(inputNums).
		Map(
			func(n interface{}) interface{} {
				return n.(int) * n.(int)
			}, 2).
		ForEach(
			func(n interface{}) {
				sum = sum + n.(int)
			})

	return sum
}
