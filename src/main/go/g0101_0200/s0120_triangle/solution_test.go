package s0120_triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	assert.Equal(t, 11, minimumTotal(triangle))
}

func TestMinimumTotal2(t *testing.T) {
	triangle := [][]int{
		{-10},
	}
	assert.Equal(t, -10, minimumTotal(triangle))
}

func TestMinimumTotal3(t *testing.T) {
	var triangle [][]int
	assert.Equal(t, 0, minimumTotal(triangle))
}
