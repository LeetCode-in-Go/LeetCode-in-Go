package s0452_minimum_number_of_arrows_to_burst_balloons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	assert.Equal(t, 2, findMinArrowShots(points))
}

func TestFindMinArrowShots2(t *testing.T) {
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	assert.Equal(t, 4, findMinArrowShots(points))
}

func TestFindMinArrowShots3(t *testing.T) {
	points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	assert.Equal(t, 2, findMinArrowShots(points))
}
