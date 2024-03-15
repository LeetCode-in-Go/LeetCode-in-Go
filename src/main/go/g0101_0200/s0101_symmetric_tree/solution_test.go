package s0101_symmetric_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetricTree(t *testing.T) {
	root := constructBinaryTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	assert.Equal(t, true, isSymmetric(root))
}

func TestSymmetricTree2(t *testing.T) {
	root := constructBinaryTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	assert.Equal(t, false, isSymmetric(root))
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
