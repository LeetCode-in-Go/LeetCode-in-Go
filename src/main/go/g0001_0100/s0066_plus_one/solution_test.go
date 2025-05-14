package s0066_plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
}

func TestPlusOne2(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
}

func TestPlusOne3(t *testing.T) {
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
}
