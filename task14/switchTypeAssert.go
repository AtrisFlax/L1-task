package main

import "fmt"

/*Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной
типа interface{}.
*/

func main() {
	do(21)
	do("hello")
	do(true)
	do(make(chan int))
	do(1.0)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%s is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("bool %v\n", v)
	case chan int:
		fmt.Printf("chan int %v \n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
