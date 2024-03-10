package s0647_palindromic_substrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, countSubstrings("abc"))
}

func TestCountSubstrings2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(6, countSubstrings("aaa"))
}
