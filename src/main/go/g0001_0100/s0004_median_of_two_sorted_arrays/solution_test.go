package s0004_median_of_two_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func TestFindMedianSortedArrays2(t *testing.T) {
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
