package s0031_next_permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	array := []int{1, 2, 3}
	nextPermutation(array)
	assert.Equal(t, []int{1, 3, 2}, array)
}

func TestNextPermutation2(t *testing.T) {
	array := []int{3, 2, 1}
	nextPermutation(array)
	assert.Equal(t, []int{1, 2, 3}, array)
}

func TestNextPermutation3(t *testing.T) {
	array := []int{1, 1, 5}
	nextPermutation(array)
	assert.Equal(t, []int{1, 5, 1}, array)
}
