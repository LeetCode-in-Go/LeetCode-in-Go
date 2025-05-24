package s0290_word_pattern

import (
	"github.com/stretchr/testify/assert"
    "testing"
)

func TestWordPattern(t *testing.T) {
	assert.True(t, wordPattern("abba", "dog cat cat dog"))
}

func TestWordPattern2(t *testing.T) {
	assert.False(t, wordPattern("abba", "dog cat cat fish"))
}

func TestWordPattern3(t *testing.T) {
	assert.False(t, wordPattern("aaaa", "dog cat cat dog"))
}

func TestWordPattern4(t *testing.T) {
	assert.False(t, wordPattern("abba", "dog dog dog dog"))
}
