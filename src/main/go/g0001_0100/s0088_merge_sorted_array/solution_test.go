package s0088_merge_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}

func TestMerge2(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	merge(nums1, 1, nums2, 0)
	assert.Equal(t, []int{1}, nums1)
}

func TestMerge3(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	assert.Equal(t, []int{1}, nums1)
}
