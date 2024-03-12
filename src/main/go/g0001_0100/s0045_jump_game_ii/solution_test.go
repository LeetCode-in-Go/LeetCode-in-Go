package s0045_jump_game_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
}

func TestJump2(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}
