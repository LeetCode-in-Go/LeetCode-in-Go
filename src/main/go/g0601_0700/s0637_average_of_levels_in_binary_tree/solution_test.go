package s0637_average_of_levels_in_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	expected := []float64{3.0, 14.5, 11.0}
	assert.Equal(t, expected, averageOfLevels(root))
}

func TestAverageOfLevels2(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 20,
		},
	}
	expected := []float64{3.0, 14.5, 11.0}
	assert.Equal(t, expected, averageOfLevels(root))
}

func TestAverageOfLevels3(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	expected := []float64{1.0}
	assert.Equal(t, expected, averageOfLevels(root))
}
