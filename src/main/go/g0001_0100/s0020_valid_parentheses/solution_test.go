package s0020_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, isValid("()"))
}

func TestIsValid2(t *testing.T) {
	assert.True(t, isValid("()[]{}"))
}

func TestIsValid3(t *testing.T) {
	assert.False(t, isValid("(]"))
}

func TestIsValid4(t *testing.T) {
	assert.False(t, isValid("([)]"))
}

func TestIsValid5(t *testing.T) {
	assert.True(t, isValid("{[]}"))
}
