package squares

import "sort"

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
квадраты в stdout.
*/

func Calculate(inputNums []int) []int {
	inputChan := gen(inputNums)
	out := squares(inputChan)

	result := make([]int, 0, len(out))
	for i := range out {
		result = append(result, i)
	}
	sort.Ints(result)
	return result
}

func squares(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- square(n)
		}
		close(out)
	}()
	return out
}

func square(n int) int {
	return n * n
}

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
