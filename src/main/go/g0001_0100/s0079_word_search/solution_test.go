package s0079_word_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	assert.True(t, exist(board, word))
}

func TestExist2(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	assert.True(t, exist(board, word))
}

func TestExis3(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	assert.False(t, exist(board, word))
}
