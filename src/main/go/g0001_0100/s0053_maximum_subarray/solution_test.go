package s0053_maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	expected := 6
	actual := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	assert.Equal(t, expected, actual)
}

func TestMaxSubArray2(t *testing.T) {
	expected := 1
	actual := maxSubArray([]int{1})
	assert.Equal(t, expected, actual)
}

func TestMaxSubArray3(t *testing.T) {
	expected := 23
	actual := maxSubArray([]int{5, 4, -1, 7, 8})
	assert.Equal(t, expected, actual)
}
