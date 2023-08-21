package main

import "fmt"

//  Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от
//  родительской структуры Human (аналог наследования).

// Human struct
type Human struct {
	intField    int
	stringField string
}

// Action embedding
type Action struct {
	Human
}

func main() {
	action := Action{}
	action.SetIntField(42)
	action.SetStringField("string")

	fmt.Println(action.IntField())
	fmt.Println(action.StringField())
}

func (h *Human) IntField() int {
	fmt.Println("IntField")
	return h.intField
}

func (h *Human) SetIntField(someField1 int) {
	fmt.Println("SetIntField")
	h.intField = someField1
}

func (h *Human) StringField() string {
	fmt.Println("StringField")
	return h.stringField
}

func (h *Human) SetStringField(someField2 string) {
	fmt.Println("SetStringField")
	h.stringField = someField2
}
