package s0103_binary_tree_zigzag_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	expected := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestZigzagLevelOrder2(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := [][]int{{1}}
	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestZigzagLevelOrder3(t *testing.T) {
	var root *TreeNode
	var expected [][]int
	assert.Equal(t, expected, zigzagLevelOrder(root))
}
