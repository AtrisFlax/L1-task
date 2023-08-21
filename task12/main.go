package main

import "fmt"

type HashSet map[string]struct{}

func (s HashSet) add(element string) {
	s[element] = struct{}{}
}

func (s HashSet) remove(element string) {
	delete(s, element)
}

func (s HashSet) has(element string) bool {
	_, ok := s[element]
	return ok
}

func main() {
	set := HashSet{}

	set.add("test1")
	set.add("test2")
	set.add("test3")

	set.add("test1")
	set.add("test1")
	set.add("test1")
	set.add("test1")

	fmt.Println(set)
	set.remove("test2")

	fmt.Println(set.has("test1"))
	fmt.Println(set.has("test2"))
	fmt.Println(set.has("test3"))
}
