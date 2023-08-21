package task11

/*
Реализовать пересечение двух неупорядоченных множеств
*/

// SimpleIntersection : Compare each element in A to each in B (O(n^2))
func SimpleIntersection[T comparable](a []T, b []T) []T {
	var result []T

	for _, v := range a {
		if contains(b, v) {
			result = append(result, v)
		}
	}

	return result
}

// HashIntersection : Put them into a hash table (O(n))
func HashIntersection[T comparable](a []T, b []T) []T {
	var result []T
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func contains[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
