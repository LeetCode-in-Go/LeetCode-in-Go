package s0013_roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
}

func TestRomanToInt2(t *testing.T) {
	assert.Equal(t, 58, romanToInt("LVIII"))
}

func TestRomanToInt3(t *testing.T) {
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
