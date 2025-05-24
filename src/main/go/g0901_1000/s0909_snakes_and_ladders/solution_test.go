package s0909_snakes_and_ladders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakesAndLadders(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}
	assert.Equal(t, 4, snakesAndLadders(board))
}

func TestSnakesAndLadders2(t *testing.T) {
	board := [][]int{
		{-1, -1},
		{-1, 3},
	}
	assert.Equal(t, 1, snakesAndLadders(board))
}

func TestSnakesAndLadders3(t *testing.T) {
	board := [][]int{
		{-1, -1, -1},
		{-1, 9, 8},
		{-1, 8, 9},
	}
	assert.Equal(t, 1, snakesAndLadders(board))
}
