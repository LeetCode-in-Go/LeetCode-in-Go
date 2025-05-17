package s0063_unique_paths_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	assert.Equal(t, 2, uniquePathsWithObstacles(obstacleGrid))
}

func TestUniquePathsWithObstacles2(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 1},
		{0, 0},
	}
	assert.Equal(t, 1, uniquePathsWithObstacles(obstacleGrid))
}
