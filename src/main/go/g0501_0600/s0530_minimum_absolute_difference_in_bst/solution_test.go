package s0530_minimum_absolute_difference_in_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	// Test case 1: [4,2,6,1,3]
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	assert.Equal(t, 1, getMinimumDifference(root1))

	// Test case 2: [1,0,48,null,null,12,49]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}
	assert.Equal(t, 1, getMinimumDifference(root2))

	// Test case 3: Single node
	root3 := &TreeNode{
		Val: 1,
	}
	assert.Equal(t, 1<<31-1, getMinimumDifference(root3))

	// Test case 4: [236,104,701,null,227,null,911]
	root4 := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}
	assert.Equal(t, 9, getMinimumDifference(root4))
}
