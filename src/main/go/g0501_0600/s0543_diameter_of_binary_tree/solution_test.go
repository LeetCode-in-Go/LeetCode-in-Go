package s0543_diameter_of_binary_tree

import (
	"reflect"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	result := diameterOfBinaryTree(createTreeNode([]int{1, 2, 3, 4, 5}))
	expected := 3
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}

func TestDiameterOfBinaryTree2(t *testing.T) {
	result := diameterOfBinaryTree(createTreeNode([]int{1, 2}))
	expected := 1
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
		nodes[i] = &TreeNode{Val: v}
	}
	for i := range nodes {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		if leftIndex < len(nodes) {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < len(nodes) {
			nodes[i].Right = nodes[rightIndex]
		}
	}
	return nodes[0]
}
