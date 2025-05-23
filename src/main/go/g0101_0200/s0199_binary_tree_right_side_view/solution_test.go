package s0199_binary_tree_right_side_view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRightSideView(t *testing.T) {
	left := &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 5},
	}
	right := &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 4},
	}
	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
	assert.Equal(t, []int{1, 3, 4}, rightSideView(root))
}

func TestRightSideView2(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(t, []int{1, 3}, rightSideView(root))
}
