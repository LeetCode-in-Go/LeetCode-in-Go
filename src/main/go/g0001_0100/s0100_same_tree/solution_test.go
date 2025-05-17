package s0100_same_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	assert.True(t, isSameTree(p, q))
}

func TestIsSameTree2(t *testing.T) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}

	q := &TreeNode{Val: 1}
	q.Right = &TreeNode{Val: 2}

	assert.False(t, isSameTree(p, q))
}

func TestIsSameTree3(t *testing.T) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 1}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 1}
	q.Right = &TreeNode{Val: 2}

	assert.False(t, isSameTree(p, q))
}
