package s0080_remove_duplicates_from_sorted_array_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, nums)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates(nums)
	assert.Equal(t, 7, k)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3, 3, 3}, nums)
}
