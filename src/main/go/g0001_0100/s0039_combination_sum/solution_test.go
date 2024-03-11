package s0039_combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	expected := [][]int{{7}, {3, 2, 2}}
	actual := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, expected, actual)
}

func TestCombinationSum2(t *testing.T) {
	expected := [][]int{{5, 3}, {3, 3, 2}, {2, 2, 2, 2}}
	actual := combinationSum([]int{2, 3, 5}, 8)
	assert.Equal(t, expected, actual)
}

func TestCombinationSum3(t *testing.T) {
	expected := [][]int(nil)
	actual := combinationSum([]int{2}, 1)
	assert.Equal(t, expected, actual)
}
