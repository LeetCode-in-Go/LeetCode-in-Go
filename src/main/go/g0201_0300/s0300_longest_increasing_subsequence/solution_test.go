package s0300_longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	result := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	expected := 4
	assert.Equal(t, expected, result)
}

func TestLengthOfLIS2(t *testing.T) {
	result := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	expected := 4
	assert.Equal(t, expected, result)
}

func TestLengthOfLIS3(t *testing.T) {
	result := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	expected := 1
	assert.Equal(t, expected, result)
}
