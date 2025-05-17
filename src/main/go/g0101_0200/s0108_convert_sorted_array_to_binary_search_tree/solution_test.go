package s0108_convert_sorted_array_to_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)

	assert.Equal(t, 0, root.Val)
	assert.Equal(t, -10, root.Left.Val)
	assert.Equal(t, 5, root.Right.Val)
	assert.Equal(t, -3, root.Left.Right.Val)
	assert.Equal(t, 9, root.Right.Right.Val)
}

func TestSortedArrayToBST2(t *testing.T) {
	nums := []int{1, 3}
	root := sortedArrayToBST(nums)

	assert.Equal(t, 1, root.Val)
	assert.Nil(t, root.Left)
	assert.Equal(t, 3, root.Right.Val)
}

func TestSortedArrayToBST3(t *testing.T) {
	var nums []int
	root := sortedArrayToBST(nums)
	assert.Nil(t, root)
}
