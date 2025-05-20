package s0127_word_ladder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	assert.Equal(t, 5, ladderLength(beginWord, endWord, wordList))
}

func TestLadderLength2(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	assert.Equal(t, 0, ladderLength(beginWord, endWord, wordList))
}

func TestLadderLength3(t *testing.T) {
	beginWord := "hit"
	endWord := "hit"
	wordList := []string{"hot"}
	assert.Equal(t, 0, ladderLength(beginWord, endWord, wordList))
}
