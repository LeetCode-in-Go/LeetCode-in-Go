package s0222_count_complete_tree_nodes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNodes(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	assert.Equal(t, 6, countNodes(root))
}

func TestCountNodes2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, 5, countNodes(root))
}

func TestCountNodes3(t *testing.T) {
	assert.Equal(t, 0, countNodes(nil))
}
