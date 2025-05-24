package s0392_is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
