package sorting_test

import (
	"go-sandbox/sorting"
	"math/rand"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := make([]int, 10)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	sorting.SelectionSort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
