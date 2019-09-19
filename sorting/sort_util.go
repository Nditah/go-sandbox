package sorting

func SelectionSort(values []int) {
	for i := 0; i < len(values); i++ {
		maxIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[maxIndex] {
				maxIndex = j
			}
		}
		tmp := values[maxIndex]
		values[maxIndex] = values[i]
		values[i] = tmp
	}
}
