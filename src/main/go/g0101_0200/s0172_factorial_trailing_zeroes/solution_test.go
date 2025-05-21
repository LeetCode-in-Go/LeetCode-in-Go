package s0172_factorial_trailing_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(3))
}

func TestTrailingZeroes2(t *testing.T) {
	assert.Equal(t, 1, trailingZeroes(5))
}

func TestTrailingZeroes3(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(0))
}
