package s0560_subarray_sum_equals_k

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
}

func TestSubarraySum2(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3))
}
