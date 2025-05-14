package s0068_text_justification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	expected := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	assert.Equal(t, expected, fullJustify(words, maxWidth))
}

func TestFullJustify2(t *testing.T) {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	expected := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	assert.Equal(t, expected, fullJustify(words, maxWidth))
}
