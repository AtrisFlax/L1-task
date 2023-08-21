package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {

	a := 1
	b := 0
	fmt.Println(swap(a, b))

	x, y := 10, 20
	x, y = y, x
	fmt.Println(x, y)
}

func swap[T any](x1 T, x2 T) (T, T) {
	return x2, x1
}
