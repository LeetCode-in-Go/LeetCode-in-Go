package s0074_search_a_2d_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	assert.Equal(t, true, searchMatrix(input, 3))
}

func TestSearchMatrix2(t *testing.T) {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	assert.Equal(t, false, searchMatrix(input, 13))
}
