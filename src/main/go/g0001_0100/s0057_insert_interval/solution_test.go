package s0057_insert_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	assert.Equal(t, [][]int{{1, 5}, {6, 9}}, insert(intervals, newInterval))
}

func TestInsert2(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	assert.Equal(t, [][]int{{1, 2}, {3, 10}, {12, 16}}, insert(intervals, newInterval))
}

func TestInsert3(t *testing.T) {
	intervals := [][]int{{1, 5}}
	newInterval := []int{2, 3}
	assert.Equal(t, [][]int{{1, 5}}, insert(intervals, newInterval))
}
