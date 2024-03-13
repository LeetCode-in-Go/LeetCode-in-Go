package s0056_merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	actual := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	assert.Equal(t, expected, actual)
}

func TestMerge2(t *testing.T) {
	expected := [][]int{{1, 5}}
	actual := merge([][]int{{1, 4}, {4, 5}})
	assert.Equal(t, expected, actual)
}
