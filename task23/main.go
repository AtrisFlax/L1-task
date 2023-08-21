package task23

//Удалить i-ый элемент из слайса

func RemoveOrdered(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func RemoveUnordered(slice []int, idx int) []int {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
