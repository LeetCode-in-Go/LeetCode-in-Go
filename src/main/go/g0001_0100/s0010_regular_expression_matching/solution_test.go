package s0010_regular_expression_matching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, false, isMatch("aa", "a"))
}

func TestIsMatch2(t *testing.T) {
	assert.Equal(t, true, isMatch("aa", "a*"))
}

func TestIsMatch3(t *testing.T) {
	assert.Equal(t, true, isMatch("ab", ".*"))
}

func TestIsMatch4(t *testing.T) {
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
}

func TestIsMatch5(t *testing.T) {
	assert.Equal(t, false, isMatch("mississippi", "mis*is*p*."))
}
