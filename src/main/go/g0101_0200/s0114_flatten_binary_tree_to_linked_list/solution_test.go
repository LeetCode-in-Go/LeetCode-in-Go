package s0114_flatten_binary_tree_to_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(root)
	expected := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}}
	assert.Equal(t, expected, root, "The flattened tree does not match the expected output.")
}

func TestFlatten2(t *testing.T) {
	root := &TreeNode{Val: 0}
	flatten(root)
	expected := &TreeNode{Val: 0}
	assert.Equal(t, expected, root, "The flattened tree does not match the expected output.")
}
