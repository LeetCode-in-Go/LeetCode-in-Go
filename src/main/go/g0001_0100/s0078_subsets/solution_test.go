package s0078_subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	expected := [][]int{(nil), {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	actual := subsets([]int{1, 2, 3})
	assert.Equal(t, expected, actual)
}

func TestSubsets2(t *testing.T) {
	expected := [][]int{(nil), {0}}
	actual := subsets([]int{0})
	assert.Equal(t, expected, actual)
}
