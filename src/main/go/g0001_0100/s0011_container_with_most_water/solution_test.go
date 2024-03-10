package s0011_container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func TestMaxArea2(t *testing.T) {
	assert.Equal(t, 1, maxArea([]int{1, 1}))
}
