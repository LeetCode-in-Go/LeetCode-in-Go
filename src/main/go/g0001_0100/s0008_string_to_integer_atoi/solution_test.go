package s0008_string_to_integer_atoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
}

func TestMyAtoi2(t *testing.T) {
	assert.Equal(t, -42, myAtoi("   -42"))
}

func TestMyAtoi3(t *testing.T) {
	assert.Equal(t, 4193, myAtoi("4193 with words"))
}

func TestMyAtoi4(t *testing.T) {
	assert.Equal(t, 0, myAtoi("words and 987"))
}

func TestMyAtoi5(t *testing.T) {
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
}
