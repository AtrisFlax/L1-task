package task20

import "strings"

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

*/

func ReverseString(input string) string {
	words := strings.Split(input, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
