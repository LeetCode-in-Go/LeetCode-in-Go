package s0073_set_matrix_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	assert.Equal(t, expected, matrix, "The matrix should be set to zeroes as expected")
}

func TestSetZeroes2(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	expected := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
	assert.Equal(t, expected, matrix, "The matrix should be set to zeroes as expected")
}
