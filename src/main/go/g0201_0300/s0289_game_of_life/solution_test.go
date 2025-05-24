package s0289_game_of_life

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	expected := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	gameOfLife(board)
	assert.Equal(t, expected, board)
}

func TestGameOfLife2(t *testing.T) {
	board := [][]int{
		{1, 1},
		{1, 0},
	}
	expected := [][]int{
		{1, 1},
		{1, 1},
	}
	gameOfLife(board)
	assert.Equal(t, expected, board)
}

func TestGameOfLife3(t *testing.T) {
	board := [][]int{
		{0, 0},
		{0, 0},
	}
	expected := [][]int{
		{0, 0},
		{0, 0},
	}
	gameOfLife(board)
	assert.Equal(t, expected, board)
}
