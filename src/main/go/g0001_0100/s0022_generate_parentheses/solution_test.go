package s0022_generate_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	)

func TestGenerateParenthesis(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	result := generateParenthesis(3)
	assert.Equal(t, expected, result, "Generated parentheses should match expected values")
}

func TestGenerateParenthesis2(t *testing.T) {
	expected := []string{"()"}
	result := generateParenthesis(1)
	assert.Equal(t, expected, result, "Generated parentheses should match expected values")
}
