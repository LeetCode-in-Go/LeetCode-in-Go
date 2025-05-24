package s0530_minimum_absolute_difference_in_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Constructs a binary tree from a level-order slice where nil means no node
func constructBinaryTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	child := 1
	for i := range nodes {
		if nodes[i] != nil && child < len(vals) {
			if child < len(nodes) && nodes[child] != nil {
				nodes[i].Left = nodes[child]
			}
			child++
			if child < len(nodes) && nodes[child] != nil {
				nodes[i].Right = nodes[child]
			}
			child++
		}
	}
	return nodes[0]
}

func TestGetMinimumDifference(t *testing.T) {
	tree := constructBinaryTree([]any{4, 2, 6, 1, 3})
	assert.Equal(t, 1, getMinimumDifference(tree))
}

func TestGetMinimumDifference2(t *testing.T) {
	tree := constructBinaryTree([]any{1, 0, 48, nil, nil, 12, 49})
	assert.Equal(t, 1, getMinimumDifference(tree))
}
