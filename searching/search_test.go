package searching_test

import (
	"github.com/andrei-punko/go-sandbox/searching"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 13, 44, 50, 51, 52}

	assert.Equal(t, searching.BinarySearch(data, -2), -1)
	assert.Equal(t, searching.BinarySearch(data, 2), 1)
	assert.Equal(t, searching.BinarySearch(data, 44), 3)
	assert.Equal(t, searching.BinarySearch(data, 45), -1)
	assert.Equal(t, searching.BinarySearch(data, 55), -1)
}

func BenchmarkBinarySearch(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data = append(data, i)
	}

	for n := 0; n < b.N; n++ {
		searching.BinarySearch(data, 500)
	}
}
