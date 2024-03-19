package s0104_maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepthBinaryTree(t *testing.T) {
	root := constructBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, 3, maxDepth(root))
}

func TestMaxDepthBinaryTree2(t *testing.T) {
	root := constructBinaryTree([]interface{}{1, nil, 2})
	assert.Equal(t, 2, maxDepth(root))
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
