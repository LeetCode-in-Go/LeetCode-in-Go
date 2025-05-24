package s0373_find_k_pairs_with_smallest_sums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	expected := [][]int{{1, 2}, {1, 4}, {1, 6}}
	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}

func TestKSmallestPairs2(t *testing.T) {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	k := 2
	expected := [][]int{{1, 1}, {1, 1}}
	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}

func TestKSmallestPairs3(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3}
	k := 3
	expected := [][]int{{1, 3}, {2, 3}}
	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}
