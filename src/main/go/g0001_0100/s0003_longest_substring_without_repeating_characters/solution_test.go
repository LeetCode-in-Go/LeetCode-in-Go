package s0003_longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
