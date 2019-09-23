package sorting_test

import (
	"go-sandbox/sorting"
	"math/rand"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	sorting.SelectionSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestBubbleSort(t *testing.T) {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	sorting.BubbleSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestCountSort(t *testing.T) {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 5
	}

	sorting.CountSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestMergeSort(t *testing.T) {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	data = sorting.MergeSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
