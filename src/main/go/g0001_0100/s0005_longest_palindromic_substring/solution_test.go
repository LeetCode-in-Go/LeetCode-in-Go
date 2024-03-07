package s0005_longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
}

func TestLongestPalindrome2(t *testing.T) {
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
