package s0427_construct_quad_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstruct(t *testing.T) {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	root := construct(grid)
	assert.False(t, root.IsLeaf)
	assert.True(t, root.Val)
	assert.True(t, root.TopLeft.IsLeaf)
	assert.False(t, root.TopLeft.Val)
	assert.True(t, root.TopRight.IsLeaf)
	assert.True(t, root.TopRight.Val)
	assert.True(t, root.BottomLeft.IsLeaf)
	assert.True(t, root.BottomLeft.Val)
	assert.True(t, root.BottomRight.IsLeaf)
	assert.False(t, root.BottomRight.Val)
}

func TestConstruct2(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	root := construct(grid)
	assert.False(t, root.IsLeaf)
	assert.True(t, root.Val)
	assert.True(t, root.TopLeft.IsLeaf)
	assert.True(t, root.TopLeft.Val)
	assert.False(t, root.TopRight.IsLeaf)
	assert.True(t, root.TopRight.Val)
	assert.True(t, root.BottomLeft.IsLeaf)
	assert.True(t, root.BottomLeft.Val)
	assert.True(t, root.BottomRight.IsLeaf)
	assert.False(t, root.BottomRight.Val)
}

func TestConstruct3(t *testing.T) {
	grid := [][]int{
		{1, 1},
		{1, 1},
	}
	root := construct(grid)
	assert.True(t, root.IsLeaf)
	assert.True(t, root.Val)
	assert.Nil(t, root.TopLeft)
	assert.Nil(t, root.TopRight)
	assert.Nil(t, root.BottomLeft)
	assert.Nil(t, root.BottomRight)
}
