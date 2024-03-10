package s1143_longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence1(t *testing.T) {
	result := longestCommonSubsequence("abcde", "ace")
	assert.Equal(t, 3, result)
}

func TestLongestCommonSubsequence2(t *testing.T) {
	result := longestCommonSubsequence("abc", "abc")
	assert.Equal(t, 3, result)
}

func TestLongestCommonSubsequence3(t *testing.T) {
	result := longestCommonSubsequence("abc", "def")
	assert.Equal(t, 0, result)
}
