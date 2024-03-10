package s0003_longest_substring_without_repeating_characters

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
