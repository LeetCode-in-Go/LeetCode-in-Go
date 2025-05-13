package s0058_length_of_last_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
}

func TestLengthOfLastWord2(t *testing.T) {
	assert.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
}

func TestLengthOfLastWord3(t *testing.T) {
	assert.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
}
