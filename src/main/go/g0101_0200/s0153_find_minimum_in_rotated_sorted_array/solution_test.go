package s0153_find_minimum_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
}

func TestFindMin2(t *testing.T) {
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func TestFindMin3(t *testing.T) {
	assert.Equal(t, 11, findMin([]int{11, 13, 15, 17}))
}
