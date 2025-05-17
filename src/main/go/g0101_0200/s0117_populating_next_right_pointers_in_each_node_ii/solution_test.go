package s0117_populating_next_right_pointers_in_each_node_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}

	result := connect(root)

	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 3, result.Left.Next.Val)
	assert.Equal(t, 7, result.Left.Right.Next.Val)
	assert.Nil(t, result.Right.Next)
}

func TestConnect2(t *testing.T) {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	result := connect(root)

	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 3, result.Left.Next.Val)
	assert.Equal(t, 6, result.Left.Right.Next.Val)
	assert.Equal(t, 7, result.Right.Left.Next.Val)
	assert.Nil(t, result.Right.Next)
}

func TestConnect3(t *testing.T) {
	var root *Node
	result := connect(root)
	assert.Nil(t, result)
}

func TestConnect4(t *testing.T) {
	root := &Node{Val: 1}
	result := connect(root)
	assert.Equal(t, 1, result.Val)
	assert.Nil(t, result.Next)
}
