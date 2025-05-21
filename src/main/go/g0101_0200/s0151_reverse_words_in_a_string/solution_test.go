package s0151_reverse_words_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))
}

func TestReverseWords2(t *testing.T) {
	assert.Equal(t, "world hello", reverseWords("  hello world  "))
}

func TestReverseWords3(t *testing.T) {
	assert.Equal(t, "example good a", reverseWords("a good   example"))
}
