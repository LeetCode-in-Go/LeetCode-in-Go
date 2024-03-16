package s0394_decode_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeString(t *testing.T) {
	result := decodeString("3[a]2[bc]")
	expected := "aaabcbc"
	assert.Equal(t, expected, result)
}

func TestDecodeString2(t *testing.T) {
	result := decodeString("3[a2[c]]")
	expected := "accaccacc"
	assert.Equal(t, expected, result)
}

func TestDecodeString3(t *testing.T) {
	result := decodeString("2[abc]3[cd]ef")
	expected := "abcabccdcdcdef"
	assert.Equal(t, expected, result)
}

func TestDecodeString4(t *testing.T) {
	result := decodeString("abc3[cd]xyz")
	expected := "abccdcdcdxyz"
	assert.Equal(t, expected, result)
}
