package squaresSum

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных
вычислений.
*/

func Calculate(inputNums []int) int {
	inputChan := gen(inputNums)
	out := squares(inputChan)

	sum := 0
	for i := range out {
		sum = sum + i
	}
	return sum
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
