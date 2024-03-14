package s0064_minimum_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	expected := 7
	actual := minPathSum(grid)
	assert.Equal(t, expected, actual)
}

func TestMinPathSum2(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	expected := 12
	actual := minPathSum(grid)
	assert.Equal(t, expected, actual)
}
