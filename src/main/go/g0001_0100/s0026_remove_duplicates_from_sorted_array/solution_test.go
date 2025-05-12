package s0026_remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{1, 2, 2}, nums)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, nums)
}
