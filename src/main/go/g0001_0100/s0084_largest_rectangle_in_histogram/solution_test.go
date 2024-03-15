package s0084_largest_rectangle_in_histogram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func TestLargestRectangleArea2(t *testing.T) {
	assert.Equal(t, 4, largestRectangleArea([]int{2, 4}))
}
