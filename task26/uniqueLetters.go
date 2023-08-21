package task26

import "strings"

// sample := "⌘こんにちは"

func uniqueLetters(input string) bool {
	strInLowerCase := strings.ToLower(input)
	alphabetSpecificSize := 33
	m := make(map[rune]*struct{}, alphabetSpecificSize)
	for _, letter := range strInLowerCase {
		_, haveValue := m[letter]
		if haveValue {
			return false
		}
		m[letter] = &struct{}{}
	}
	return true
}
