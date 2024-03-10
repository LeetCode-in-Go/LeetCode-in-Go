package s0007_reverse_integer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
}

func TestReverse2(t *testing.T) {
	assert.Equal(t, -321, reverse(-123))
}

func TestReverse3(t *testing.T) {
	assert.Equal(t, 21, reverse(120))
}
