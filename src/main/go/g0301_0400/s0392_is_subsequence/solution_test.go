package s0392_is_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	assert.True(t, isSubsequence("abc", "ahbgdc"))
}

func TestIsSubsequence2(t *testing.T) {
	assert.False(t, isSubsequence("axc", "ahbgdc"))
}

func TestIsSubsequence3(t *testing.T) {
	assert.True(t, isSubsequence("", "ahbgdc"))
}

func TestIsSubsequence4(t *testing.T) {
	assert.False(t, isSubsequence("abc", ""))
}
