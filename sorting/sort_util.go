package sorting

func SelectionSort(values []int) {
	for i := 0; i < len(values); i++ {
		minIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		values[i], values[minIndex] = values[minIndex], values[i]
	}
}

func InsertionSort(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			} else {
				break
			}
		}
	}
}

func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		swapHappened := false
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapHappened = true
			}
		}
		if !swapHappened {
			return
		}
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
	mergeSort(values, len(values))
	return values
}

func mergeSort(values []int, n int) {
	if n < 2 {
		return
	}
	mid := n / 2
	l := make([]int, mid)
	r := make([]int, n-mid)

	for i := 0; i < mid; i++ {
		l[i] = values[i]
	}
	for i := mid; i < n; i++ {
		r[i-mid] = values[i]
	}
	mergeSort(l, mid)
	mergeSort(r, n-mid)

	merge(values, l, r, mid, n-mid)
}

func merge(items []int, l []int, r []int, left int, right int) {
	i := 0
	j := 0
	k := 0
	for i < left && j < right {
		if l[i] <= r[j] {
			items[k] = l[i]
			k++
			i++
		} else {
			items[k] = r[j]
			k++
			j++
		}
	}
	for i < left {
		items[k] = l[i]
		k++
		i++
	}
	for j < right {
		items[k] = r[j]
		k++
		j++
	}
}

func QuickSort(values []int) []int {
	quickSort(values, 0, len(values)-1)
	return values
}

func quickSort(values []int, l int, r int) {
	if l < r {
		q := partition(values, l, r)
		quickSort(values, l, q)
		quickSort(values, q+1, r)
	}
}

func partition(values []int, l int, r int) int {
	center := (l + r) / 2
	centerValue := values[center]
	i := l
	j := r
	for i <= j {
		for values[i] < centerValue {
			i++
		}
		for values[j] > centerValue {
			j--
		}
		if i >= j {
			break
		}
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}
	return j
}
