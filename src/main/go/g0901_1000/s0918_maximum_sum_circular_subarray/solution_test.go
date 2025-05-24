package s0918_maximum_sum_circular_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	assert.Equal(t, 3, maxSubarraySumCircular([]int{1, -2, 3, -2}))
}

func TestMaxSubarraySumCircular2(t *testing.T) {
	assert.Equal(t, 10, maxSubarraySumCircular([]int{5, -3, 5}))
}

func TestMaxSubarraySumCircular3(t *testing.T) {
	assert.Equal(t, -2, maxSubarraySumCircular([]int{-3, -2, -3}))
}

func TestMaxSubarraySumCircular4(t *testing.T) {
	assert.Equal(t, 1, maxSubarraySumCircular([]int{1}))
}

func TestMaxSubarraySumCircular5(t *testing.T) {
	assert.Equal(t, 15, maxSubarraySumCircular([]int{1, 2, 3, 4, 5}))
}

func TestMaxSubarraySumCircular6(t *testing.T) {
	assert.Equal(t, -1, maxSubarraySumCircular([]int{-1, -2, -3, -4, -5}))
}

func TestMaxSubarraySumCircular7(t *testing.T) {
	assert.Equal(t, 12, maxSubarraySumCircular([]int{3, -1, 2, -1, 8}))
}
