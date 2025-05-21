package s0150_evaluate_reverse_polish_notation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	assert.Equal(t, 9, evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func TestEvalRPN2(t *testing.T) {
	assert.Equal(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
}

func TestEvalRPN3(t *testing.T) {
	assert.Equal(t, 22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
