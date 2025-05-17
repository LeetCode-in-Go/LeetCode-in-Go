package s0097_interleaving_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	assert.True(t, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

func TestIsInterleave2(t *testing.T) {
	assert.False(t, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}

func TestIsInterleave3(t *testing.T) {
	assert.True(t, isInterleave("", "", ""))
}

func TestIsInterleave4(t *testing.T) {
	assert.False(t, isInterleave("a", "b", "a"))
}
