package s0001_two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
}

func TwoSum2(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
}

func TwoSum3(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func TwoSum4(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, twoSum([]int{3, 3}, 7))
}
