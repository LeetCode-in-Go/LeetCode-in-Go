package s0283_move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, array)
}

func TestMoveZeroes2(t *testing.T) {
	array := []int{0}
	moveZeroes(array)
	assert.Equal(t, []int{0}, array)
}
