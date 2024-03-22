package s0189_rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(array, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, array)
}

func TestRotate2(t *testing.T) {
	array := []int{-1, -100, 3, 99}
	rotate(array, 2)
	assert.Equal(t, []int{3, 99, -1, -100}, array)
}
