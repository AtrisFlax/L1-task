package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример
реализации.
//After slicing gc wont collect entire v despite we using only first 100
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

/* strings.Clone decouple usage of huge string resource
func someFunc(hugeString *string) string {
	v := createHugeString(1 << 10)
	justString:= strings.Clone(string(v)[:100])
}
*/

//func main() {
//
//}
