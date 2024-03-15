package s0010_regular_expression_matching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
}

func TestIsMatch2(t *testing.T) {
	assert.True(t, isMatch("aa", "a*"))
}

func TestIsMatch3(t *testing.T) {
	assert.True(t, isMatch("ab", ".*"))
}

func TestIsMatch4(t *testing.T) {
	assert.True(t, isMatch("aab", "c*a*b"))
}

func TestIsMatch5(t *testing.T) {
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
}
