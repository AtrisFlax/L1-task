package squares

import (
	"github.com/jucardi/go-streams/streams"
	"sort"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
квадраты в stdout.
*/

func CalculateStreams(inputNums []int) []int {
	squares := streams.
		From(inputNums).
		Map(
			func(n interface{}) interface{} {
				return n.(int) * n.(int)
			}, 2).
		ToArray().([]int)

	sort.Ints(squares)
	return squares
}
