package s0105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructBinaryTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	actual := buildTree(preorder, inorder)
	expected := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, expected, actual, "Constructed binary tree does not match expected output.")
}

func TestConstructBinaryTree2(t *testing.T) {
	preorder := []int{-1}
	inorder := []int{-1}
	actual := buildTree(preorder, inorder)
	expected := &TreeNode{Val: -1}
	assert.Equal(t, expected, actual, "Constructed binary tree does not match expected output.")
}
