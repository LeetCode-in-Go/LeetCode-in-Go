package s0226_invert_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	root := newTreeNode(4,
		newTreeNode(2, newTreeNode(1, nil, nil), newTreeNode(3, nil, nil)),
		newTreeNode(7, newTreeNode(6, nil, nil), newTreeNode(9, nil, nil)),
	)
	expected := newTreeNode(4,
		newTreeNode(7, newTreeNode(9, nil, nil), newTreeNode(6, nil, nil)),
		newTreeNode(2, newTreeNode(3, nil, nil), newTreeNode(1, nil, nil)),
	)
	assert.Equal(t, expected, invertTree(root))
}

func TestInvertTree2(t *testing.T) {
	root := newTreeNode(2, newTreeNode(1, nil, nil), newTreeNode(3, nil, nil))
	expected := newTreeNode(2, newTreeNode(3, nil, nil), newTreeNode(1, nil, nil))
	assert.Equal(t, expected, invertTree(root))
}

func TestInvertTree3(t *testing.T) {
	root := newTreeNode(1, nil, nil)
	expected := newTreeNode(1, nil, nil)
	assert.Equal(t, expected, invertTree(root))
}

func newTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}
