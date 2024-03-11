package s0034_find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	expected := []int{3, 4}
	actual := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.Equal(t, expected, actual)
}

func TestSearchRange2(t *testing.T) {
	expected := []int{-1, -1}
	actual := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	assert.Equal(t, expected, actual)
}

func TestSearchRange3(t *testing.T) {
	expected := []int{-1, -1}
	actual := searchRange([]int{}, 0)
	assert.Equal(t, expected, actual)
}
