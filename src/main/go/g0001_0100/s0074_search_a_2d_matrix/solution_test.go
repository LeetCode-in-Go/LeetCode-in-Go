package s0074_search_a_2d_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	assert.True(t, searchMatrix(input, 3))
}

func TestSearchMatrix2(t *testing.T) {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	assert.True(t, searchMatrix(input, 13))
}
