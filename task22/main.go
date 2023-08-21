package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

func main() {
	a, _ := new(big.Int).SetString("2000000", 10)
	b, _ := new(big.Int).SetString("2000001", 10)

	result := new(big.Int)
	fmt.Println(result.Mul(a, b))
	fmt.Println(result.Div(a, b))
	fmt.Println(result.Add(a, b))
	fmt.Println(result.Sub(a, b))
}
