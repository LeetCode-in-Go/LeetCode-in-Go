package s0094_binary_tree_inorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}
	expected := []int{1, 3, 2}
	actual := inorderTraversal(root)
	assert.Equal(t, expected, actual)
}

func TestInorderTraversal2(t *testing.T) {
	var root *TreeNode
	expected := []int{}
	actual := inorderTraversal(root)
	assert.Equal(t, expected, actual)
}

func TestInorderTraversal3(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	actual := inorderTraversal(root)
	assert.Equal(t, expected, actual)
}
