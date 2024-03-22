package s0230_kth_smallest_element_in_a_bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	root := NewTreeNode([]interface{}{3, 1, 4, nil, 2})
	assert.Equal(t, 1, kthSmallest(root, 1))
}

func TestKthSmallest2(t *testing.T) {
	root := NewTreeNode([]interface{}{5, 3, 6, 2, 4, nil, nil, 1})
	assert.Equal(t, 3, kthSmallest(root, 3))
}

func NewTreeNode(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	var root *TreeNode
	nodes := make([]*TreeNode, len(values))
	for i, value := range values {
		if value != nil {
			val, ok := value.(int)
			if !ok {
				fmt.Printf("Value at index %d is not of type int\n", i)
				return nil
			}
			nodes[i] = &TreeNode{Val: val}
		}
	}

	for i, node := range nodes {
		if node != nil {
			leftIndex := 2*i + 1
			if leftIndex < len(values) && nodes[leftIndex] != nil {
				node.Left = nodes[leftIndex]
			}
			rightIndex := 2*i + 2
			if rightIndex < len(values) && nodes[rightIndex] != nil {
				node.Right = nodes[rightIndex]
			}
		}
	}

	root = nodes[0]
	return root
}
