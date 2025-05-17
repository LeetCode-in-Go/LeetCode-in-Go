package s0106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)

	assert.Equal(t, 3, root.Val)
	assert.Equal(t, 9, root.Left.Val)
	assert.Equal(t, 20, root.Right.Val)
	assert.Equal(t, 15, root.Right.Left.Val)
	assert.Equal(t, 7, root.Right.Right.Val)
}

func TestBuildTree2(t *testing.T) {
	inorder := []int{-1}
	postorder := []int{-1}
	root := buildTree(inorder, postorder)

	assert.Equal(t, -1, root.Val)
	assert.Nil(t, root.Left)
	assert.Nil(t, root.Right)
}
