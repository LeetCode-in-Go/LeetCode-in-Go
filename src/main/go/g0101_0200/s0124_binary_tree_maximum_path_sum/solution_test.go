package s0124_binary_tree_maximum_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	assert.Equal(t, 6, maxPathSum(root))
}

func TestMaxPathSum2(t *testing.T) {
	root := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, 42, maxPathSum(root))
}
