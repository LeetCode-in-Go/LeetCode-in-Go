package s0201_bitwise_and_of_numbers_range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	assert.Equal(t, 4, rangeBitwiseAnd(5, 7))
}

func TestRangeBitwiseAnd2(t *testing.T) {
	assert.Equal(t, 0, rangeBitwiseAnd(0, 0))
}

func TestRangeBitwiseAnd3(t *testing.T) {
	assert.Equal(t, 0, rangeBitwiseAnd(1, 2147483647))
}
