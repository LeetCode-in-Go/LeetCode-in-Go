package s0152_maximum_product_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	assert.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
}

func TestMaxProduct2(t *testing.T) {
	assert.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
}
