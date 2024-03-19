package s0102_binary_tree_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := constructBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	assert.Equal(t, expected, levelOrder(root))
}

func TestLevelOrder2(t *testing.T) {
	root := constructBinaryTree([]interface{}{1})
	expected := [][]int{{1}}
	assert.Equal(t, expected, levelOrder(root))
}

func TestLevelOrder3(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil))
}

func constructBinaryTree(treeValues []interface{}) *TreeNode {
	if len(treeValues) == 0 || treeValues[0] == nil {
		return nil
	}
	root := &TreeNode{Val: treeValues[0].(int)}
	queue := []*TreeNode{root}
	for i := 1; i < len(treeValues); i++ {
		curr := queue[0]
		queue = queue[1:]
		if treeValues[i] != nil {
			curr.Left = &TreeNode{Val: treeValues[i].(int)}
			queue = append(queue, curr.Left)
		}
		if i++; i < len(treeValues) && treeValues[i] != nil {
			curr.Right = &TreeNode{Val: treeValues[i].(int)}
			queue = append(queue, curr.Right)
		}
	}
	return root
}
