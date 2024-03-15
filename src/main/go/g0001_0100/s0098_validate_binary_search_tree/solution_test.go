package s0098_validate_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	treeNode1 := &TreeNode{Val: 2}
	treeNode1.Left = &TreeNode{Val: 1}
	treeNode1.Right = &TreeNode{Val: 3}
	assert.True(t, isValidBST(treeNode1), "Expected true, but got false")
}

func TestIsValidBST2(t *testing.T) {
	treeNode := &TreeNode{Val: 5}
	treeNode.Left = &TreeNode{Val: 1}
	treeNode.Right = &TreeNode{Val: 4}
	treeNode.Right.Left = &TreeNode{Val: 3}
	treeNode.Right.Right = &TreeNode{Val: 6}
	assert.False(t, isValidBST(treeNode), "Expected false, but got true")
}
