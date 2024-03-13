package s0055_jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	expected := true
	actual := canJump([]int{2, 3, 1, 1, 4})
	assert.Equal(t, expected, actual)
}

func TestCanJump2(t *testing.T) {
	expected := false
	actual := canJump([]int{3, 2, 1, 0, 4})
	assert.Equal(t, expected, actual)
}
