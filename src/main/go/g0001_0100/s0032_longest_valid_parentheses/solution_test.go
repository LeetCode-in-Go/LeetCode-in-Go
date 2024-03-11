package s0032_longest_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("(()"))
}

func TestLongestValidParentheses2(t *testing.T) {
	assert.Equal(t, 4, longestValidParentheses(")()())"))
}

func TestLongestValidParentheses3(t *testing.T) {
	assert.Equal(t, 0, longestValidParentheses(""))
}
