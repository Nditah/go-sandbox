package sorting

func SelectionSort(values []int) {
	for i := 0; i < len(values); i++ {
		maxIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[maxIndex] {
				maxIndex = j
			}
		}
		swap(values, i, maxIndex)
	}
}

func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		swapHappened := false;
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				swap(values, j, j+1)
				swapHappened = true
			}
		}
		if !swapHappened {
			return
		};
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
