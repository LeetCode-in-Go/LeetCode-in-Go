package s0129_sum_root_to_leaf_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	assert.Equal(t, 25, sumNumbers(root))
}

func TestSumNumbers2(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}
	assert.Equal(t, 1026, sumNumbers(root))
}

func TestSumNumbers3(t *testing.T) {
	root := &TreeNode{Val: 1}
	assert.Equal(t, 1, sumNumbers(root))
}
