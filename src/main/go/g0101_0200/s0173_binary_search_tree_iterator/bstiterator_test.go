package s0173_binary_search_tree_iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	left := &TreeNode{Val: 3}
	right := &TreeNode{
		Val:   15,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20},
	}
	root := &TreeNode{
		Val:   7,
		Left:  left,
		Right: right,
	}
	iterator := Constructor(root)
	assert.Equal(t, 3, iterator.Next())
	assert.Equal(t, 7, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 9, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 15, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 20, iterator.Next())
	assert.False(t, iterator.HasNext())
}
