package searching_test

import (
	"github.com/stretchr/testify/assert"
	"go-sandbox/searching"
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
