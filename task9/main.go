package main

import (
	"fmt"
	"time"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	inputValuesChannel := make(chan int)
	go func() {
		for _, n := range values {
			inputValuesChannel <- n
		}
		defer close(inputValuesChannel)
	}()

	resultChannel := make(chan int)
	go func() {
		for n := range inputValuesChannel {
			resultChannel <- multiplyByTwo(n)
		}
		defer close(resultChannel)
	}()

	for r := range resultChannel {
		fmt.Println(r)
	}

	time.Sleep(time.Second * 2)
}

func multiplyByTwo(x int) int {
	return 2 * x
}
