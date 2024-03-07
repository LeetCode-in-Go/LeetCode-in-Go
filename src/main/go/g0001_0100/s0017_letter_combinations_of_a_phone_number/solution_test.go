package s0017_letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	result := letterCombinations("23")
	assert.Equal(t, expected, result, "The two words should be the same.")
}

func TestLetterCombinations2(t *testing.T) {
	expected := []string(nil)
	result := letterCombinations("")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations3(t *testing.T) {
	expected := []string{"a", "b", "c"}
	result := letterCombinations("2")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations4(t *testing.T) {
	expected := []string{"g", "h", "i"}
	result := letterCombinations("4")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations5(t *testing.T) {
	expected := []string{"j", "k", "l"}
	result := letterCombinations("5")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations6(t *testing.T) {
	expected := []string{"m", "n", "o"}
	result := letterCombinations("6")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations7(t *testing.T) {
	expected := []string{"p", "q", "r", "s"}
	result := letterCombinations("7")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations8(t *testing.T) {
	expected := []string{"t", "u", "v"}
	result := letterCombinations("8")
	assert.Equal(t, expected, result)
}

func TestLetterCombinations9(t *testing.T) {
	expected := []string{"w", "x", "y", "z"}
	result := letterCombinations("9")
	assert.Equal(t, expected, result)
}
