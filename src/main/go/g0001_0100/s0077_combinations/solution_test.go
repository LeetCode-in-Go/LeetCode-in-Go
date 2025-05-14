package s0077_combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	expected := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	assert.Equal(t, expected, combine(4, 2))
}

func TestCombine2(t *testing.T) {
	expected := [][]int{{1}}
	assert.Equal(t, expected, combine(1, 1))
}
