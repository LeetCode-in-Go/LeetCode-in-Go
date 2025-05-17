package s0112_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	assert.True(t, hasPathSum(root, 22))
}

func TestHasPathSum2(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	assert.False(t, hasPathSum(root, 5))
}

func TestHasPathSum3(t *testing.T) {
	var root *TreeNode
	assert.False(t, hasPathSum(root, 0))
}

func TestHasPathSum4(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	assert.False(t, hasPathSum(root, 1))
}
