package s0139_word_break

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	assert.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))
}

func TestWordBreak2(t *testing.T) {
	assert.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
}

func TestWordBreak3(t *testing.T) {
	assert.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
