package s0130_surrounded_regions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	assert.Equal(t, expected, board)
}

func TestSolve2(t *testing.T) {
	board := [][]byte{
		{'X'},
	}
	expected := [][]byte{
		{'X'},
	}
	solve(board)
	assert.Equal(t, expected, board)
}

func TestSolve3(t *testing.T) {
	board := [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	expected := [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	solve(board)
	assert.Equal(t, expected, board)
}

func TestSolve4(t *testing.T) {
	var board [][]byte
	solve(board)
	assert.Nil(t, board)
}
