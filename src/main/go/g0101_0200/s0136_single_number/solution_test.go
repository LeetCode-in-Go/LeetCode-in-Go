package s0136_single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
}

func TestSingleNumber2(t *testing.T) {
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
}

func TestSingleNumber3(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{1}))
}
