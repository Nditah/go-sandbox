package sorting_test

import (
	"go-sandbox/sorting"
	"math/rand"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := prepareArray(100)
	sorting.SelectionSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestInsertionSort(t *testing.T) {
	data := prepareArray(100)
	sorting.InsertionSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestBubbleSort(t *testing.T) {
	data := prepareArray(100)
	sorting.BubbleSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestCountSort(t *testing.T) {
	const maxValue = 5
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % maxValue
	}

	sorting.CountSort(data, maxValue)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestMergeSort(t *testing.T) {
	data := prepareArray(100)
	data = sorting.MergeSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestQuickSort(t *testing.T) {
	data := prepareArray(100)
	data = sorting.QuickSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

const itemsCount = 1000

func BenchmarkBubbleSort(b *testing.B) {
	data := prepareArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.BubbleSort(data)
	}
}

func BenchmarkBubbleSortSorted(b *testing.B) {
	data := prepareSortedArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.BubbleSort(data)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	data := prepareArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.InsertionSort(data)
	}
}

func BenchmarkInsertionSortSorted(b *testing.B) {
	data := prepareSortedArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.InsertionSort(data)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	data := prepareArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.MergeSort(data)
	}
}

func BenchmarkMergeSortSorted(b *testing.B) {
	data := prepareSortedArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.MergeSort(data)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	data := prepareArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.SelectionSort(data)
	}
}

func BenchmarkSelectionSortSorted(b *testing.B) {
	data := prepareSortedArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.SelectionSort(data)
	}
}

func BenchmarkCountSort(b *testing.B) {
	const maxValue = 200
	data := make([]int, itemsCount)
	for i := range data {
		data[i] = rand.Int() % maxValue
	}

	for n := 0; n < b.N; n++ {
		sorting.CountSort(data, maxValue)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	data := prepareArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.QuickSort(data)
	}
}

func BenchmarkQuickSortSorted(b *testing.B) {
	data := prepareSortedArray(itemsCount)
	for n := 0; n < b.N; n++ {
		sorting.QuickSort(data)
	}
}

func prepareArray(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

func prepareSortedArray(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	return data
}
