package s0067_add_binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
}

func TestAddBinary2(t *testing.T) {
	assert.Equal(t, "10101", addBinary("1010", "1011"))
}
