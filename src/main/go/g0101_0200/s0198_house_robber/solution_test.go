package s0198_house_robber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
}

func TestRob2(t *testing.T) {
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
}
