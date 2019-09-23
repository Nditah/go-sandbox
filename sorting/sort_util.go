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

func InsertionSort(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				swap(values, j, j-1)
			} else {
				break
			}
		}
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

func CountSort(values []int, max int) {
	counts := make([]int, max)
	for _, value := range values {
		counts[value]++
	}

	counter := 0
	for value, count := range counts {
		for i := 0; i < count; i++ {
			values[counter] = value
			counter++
		}
	}
}

func MergeSort(values []int) []int {
	if len(values) == 1 {
		return values
	}

	center := len(values) / 2
	left := values[:center]
	right := values[center:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var result []int
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}
	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}
	return result
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
