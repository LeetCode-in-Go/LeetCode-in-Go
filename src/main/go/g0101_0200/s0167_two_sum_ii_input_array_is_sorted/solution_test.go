package s0167_two_sum_ii_input_array_is_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
}

func TestTwoSum2(t *testing.T) {
	assert.Equal(t, []int{1, 3}, twoSum([]int{2, 3, 4}, 6))
}

func TestTwoSum3(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{-1, 0}, -1))
}
