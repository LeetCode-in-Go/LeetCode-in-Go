package s0048_rotate_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(input)
	assert.Equal(t, expected, input)
}

func TestRotate2(t *testing.T) {
	expected := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}
	input := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(input)
	assert.Equal(t, expected, input)
}
