package s0075_sort_colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	array := []int{2, 0, 2, 1, 1, 0}
	sortColors(array)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, array)
}

func TestSortColors2(t *testing.T) {
	array := []int{2, 0, 1}
	sortColors(array)
	assert.Equal(t, []int{0, 1, 2}, array)
}
