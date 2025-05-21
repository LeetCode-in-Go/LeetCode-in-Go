package s0137_single_number_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 3, singleNumber([]int{2, 2, 3, 2}))
}

func TestSingleNumber2(t *testing.T) {
	assert.Equal(t, 99, singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
