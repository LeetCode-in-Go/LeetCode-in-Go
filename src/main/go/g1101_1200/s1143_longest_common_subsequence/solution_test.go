package s1143_longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence1(t *testing.T) {
	assert := assert.New(t)
	result := longestCommonSubsequence("abcde", "ace")
	assert.Equal(3, result)
}

func TestLongestCommonSubsequence2(t *testing.T) {
	assert := assert.New(t)
	result := longestCommonSubsequence("abc", "abc")
	assert.Equal(3, result)
}

func TestLongestCommonSubsequence3(t *testing.T) {
	assert := assert.New(t)
	result := longestCommonSubsequence("abc", "def")
	assert.Equal(0, result)
}
