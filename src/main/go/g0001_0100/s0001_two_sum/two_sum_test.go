package s0001_two_sum

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestTwoSum(t *testing.T) {
    assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
}

func TestTwoSum2(t *testing.T) {
    assert.Equal(t, []int{0, 2}, twoSum([]int{7, 11, 15}, 22))
}

func TestTwoSum3(t *testing.T) {
    assert.Equal(t, []int{0, 1}, twoSum([]int{11, 15}, 26))
}
