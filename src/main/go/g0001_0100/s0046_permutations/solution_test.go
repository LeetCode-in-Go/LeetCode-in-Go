package s0046_permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		permute([]int{1, 2, 3}))
}

func TestPermute2(t *testing.T) {
	assert.Equal(t, [][]int{{0, 1}, {1, 0}}, permute([]int{0, 1}))
}

func TestPermute3(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, permute([]int{1}))
}
