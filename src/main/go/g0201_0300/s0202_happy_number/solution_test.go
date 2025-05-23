package s0202_happy_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHappy(t *testing.T) {
	assert.True(t, isHappy(19))
}

func TestIsHappy2(t *testing.T) {
	assert.False(t, isHappy(2))
}
