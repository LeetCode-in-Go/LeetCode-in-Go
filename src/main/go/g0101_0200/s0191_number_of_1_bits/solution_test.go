package s0191_number_of_1_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(0b00000000000000000000000000001011))
}

func TestHammingWeight2(t *testing.T) {
	assert.Equal(t, 1, hammingWeight(0b00000000000000000000000010000000))
}
