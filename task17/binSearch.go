package binsearch

func BinarySearch(a []int, x int) int {
	switch len(a) {
	case 0:
		return -1
	case 1:
		if a[0] == x {
			return 0
		}
		return -1
	default:
		minValue, maxValue := 0, len(a)
		for minValue <= maxValue {
			middle := (minValue + maxValue) / 2

			if a[middle] == x {
				return middle
			}

			if a[middle] < x {
				minValue = middle + 1
			} else {
				maxValue = middle - 1
			}
		}
	}
	return -1
}
