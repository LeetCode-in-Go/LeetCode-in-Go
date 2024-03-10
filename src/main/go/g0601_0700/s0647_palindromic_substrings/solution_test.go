package s0647_palindromic_substrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	assert.Equal(t, 3, countSubstrings("abc"))
}

func TestCountSubstrings2(t *testing.T) {
	assert.Equal(t, 6, countSubstrings("aaa"))
}
