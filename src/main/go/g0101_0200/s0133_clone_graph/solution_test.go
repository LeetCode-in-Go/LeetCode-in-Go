package s0133_clone_graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	cloned := cloneGraph(node1)
	assert.NotNil(t, cloned)
	assert.Equal(t, 1, cloned.Val)
	assert.Equal(t, 2, len(cloned.Neighbors))
	assert.Equal(t, 2, cloned.Neighbors[0].Val)
	assert.Equal(t, 4, cloned.Neighbors[1].Val)
}

func TestCloneGraph2(t *testing.T) {
	var node *Node
	cloned := cloneGraph(node)
	assert.Nil(t, cloned)
}

func TestCloneGraph3(t *testing.T) {
	node := &Node{Val: 1}
	cloned := cloneGraph(node)
	assert.NotNil(t, cloned)
	assert.Equal(t, 1, cloned.Val)
	assert.Empty(t, cloned.Neighbors)
}
