package s0005_longest_palindromic_substring

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
}

func TestLongestPalindrome2(t *testing.T) {
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
