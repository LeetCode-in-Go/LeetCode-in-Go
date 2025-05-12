package s0012_integer_to_roman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "III", intToRoman(3))
}

func TestIntToRoman2(t *testing.T) {
	assert.Equal(t, "IV", intToRoman(4))
}

func TestIntToRoman3(t *testing.T) {
	assert.Equal(t, "IX", intToRoman(9))
}

func TestIntToRoman4(t *testing.T) {
	assert.Equal(t, "LVIII", intToRoman(58))
}

func TestIntToRoman5(t *testing.T) {
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
}
