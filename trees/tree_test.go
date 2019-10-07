package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeToListWide(t *testing.T) {
	node4 := TreeNode{value: 4}
	node5 := TreeNode{value: 5}
	node3 := TreeNode{left: &node5, value: 3, right: &node4}
	node2 := TreeNode{value: 2}
	tree := TreeNode{left: &node2, value: 1, right: &node3}

	expectedArray := []int{1, 2, 3, 5, 4}
	assert.Equal(t, []int(nil), TreeToListWide(nil), "Should return nil for nil input")
	assert.Equal(t, expectedArray, TreeToListWide(&tree))
}

func TestTreeToListDeep(t *testing.T) {
	node4 := TreeNode{value: 4}
	node5 := TreeNode{value: 5}
	node3 := TreeNode{left: &node5, value: 3, right: &node4}
	node2 := TreeNode{value: 2}
	tree := TreeNode{left: &node2, value: 1, right: &node3}

	expectedArray := []int{2, 1, 5, 3, 4}
	assert.Equal(t, []int(nil), TreeToListDeep(nil), "Should return nil for nil input")
	assert.Equal(t, expectedArray, TreeToListDeep(&tree))
}

func BenchmarkTreeToListDeep(b *testing.B) {
	tree := BuildTree(10)
	for n := 0; n < b.N; n++ {
		TreeToListDeep(&tree)
	}
}

func BenchmarkTreeToListWide(b *testing.B) {
	tree := BuildTree(10)
	for n := 0; n < b.N; n++ {
		TreeToListWide(&tree)
	}
}

func BenchmarkBuildTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildTree(10)
	}
}
