package s0560_subarray_sum_equals_k

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSubarraySum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, subarraySum([]int{1, 1, 1}, 2))
}

func TestSubarraySum2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, subarraySum([]int{1, 2, 3}, 3))
}
