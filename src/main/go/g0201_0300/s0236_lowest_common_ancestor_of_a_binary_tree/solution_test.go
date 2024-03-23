package s0236_lowest_common_ancestor_of_a_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	leftNodeLeftNode := &TreeNode{Val: 6}
	leftNodeRightNode := &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}
	leftNode := &TreeNode{Val: 5, Left: leftNodeLeftNode, Right: leftNodeRightNode}
	rightNode := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
	root := &TreeNode{Val: 3, Left: leftNode, Right: rightNode}
	assert.Equal(t, 3, lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1}).Val)
}

func TestLowestCommonAncestor2(t *testing.T) {
	leftNodeLeftNode := &TreeNode{Val: 6}
	leftNodeRightNode := &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}
	leftNode := &TreeNode{Val: 5, Left: leftNodeLeftNode, Right: leftNodeRightNode}
	rightNode := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
	root := &TreeNode{Val: 3, Left: leftNode, Right: rightNode}
	assert.Equal(t, 5, lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 4}).Val)
}

func TestLowestCommonAncestor3(t *testing.T) {
	assert.Equal(t, 2, lowestCommonAncestor(&TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, &TreeNode{Val: 2}, &TreeNode{Val: 1}).Val)
}
