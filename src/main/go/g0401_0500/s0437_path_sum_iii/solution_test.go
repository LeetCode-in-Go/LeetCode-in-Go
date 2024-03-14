package s0437_path_sum_iii

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	root := createTreeNode([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1})
	result := pathSum(root, 8)
	expected := 3
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}

func TestPathSum2(t *testing.T) {
	root := createTreeNode([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1})
	result := pathSum(root, 22)
	expected := 3
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}

func createTreeNode(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(values))
	for i, v := range values {
		if v == 0 {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v}
		}
	}
	for i := range nodes {
		if nodes[i] != nil {
			leftIndex := 2*i + 1
			rightIndex := 2*i + 2
			if leftIndex < len(nodes) {
				nodes[i].Left = nodes[leftIndex]
			}
			if rightIndex < len(nodes) {
				nodes[i].Right = nodes[rightIndex]
			}
		}
	}
	return nodes[0]
}
