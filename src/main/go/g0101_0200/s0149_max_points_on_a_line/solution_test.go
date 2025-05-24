package s0149_max_points_on_a_line

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPoints(t *testing.T) {
	input := [][]int{{1, 1}, {2, 2}, {3, 3}}
	assert.Equal(t, 3, maxPoints(input))
}

func TestMaxPoints2(t *testing.T) {
	input := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	assert.Equal(t, 4, maxPoints(input))
}
