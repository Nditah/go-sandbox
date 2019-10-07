package trees

import "math/rand"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func TreeToListWide(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}
	res := []int{tree.value}
	if tree.left != nil {
		leftList := TreeToListWide(tree.left)
		res = append(res, leftList...)
	}
	if tree.right != nil {
		rightList := TreeToListWide(tree.right)
		res = append(res, rightList...)
	}
	return res
}

func TreeToListDeep(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}
	res := make([]int, 0)
	if tree.left != nil {
		leftList := TreeToListDeep(tree.left)
		res = append(res, leftList...)
	}
	res = append(res, tree.value)
	if tree.right != nil {
		rightList := TreeToListDeep(tree.right)
		res = append(res, rightList...)
	}
	return res
}

func BuildTree(height int) TreeNode {
	rootNode := TreeNode{value: rand.Int()}
	if height == 1 {
		return rootNode
	}

	if rand.Int()%5 < 3 {
		tree := BuildTree(height - 1)
		rootNode.left = &tree
	}
	if rand.Int()%5 < 3 {
		tree := BuildTree(height - 1)
		rootNode.right = &tree
	}

	return rootNode
}
