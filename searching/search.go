package searching

func BinarySearch(values []int, value int) int {
	left := 0
	right := len(values)

	for ; right-left > 1; {
		center := (left + right) / 2

		if value < values[center] {
			right = center
		} else if values[center] == value {
			return center
		} else {
			left = center
		}
	}
	return -1
}
