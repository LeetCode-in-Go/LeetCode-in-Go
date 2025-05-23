package s0224_basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, 2, calculate("1 + 1"))
}

func TestCalculate2(t *testing.T) {
	assert.Equal(t, 3, calculate(" 2-1 + 2 "))
}

func TestCalculate3(t *testing.T) {
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
}
