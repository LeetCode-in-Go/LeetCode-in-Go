package s0212_word_search_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	expected := []string{"oath", "eat"}
	assert.Equal(t, expected, findWords(board, words))
}

func TestFindWords2(t *testing.T) {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words := []string{"abcb"}
	assert.Empty(t, findWords(board, words))
}
