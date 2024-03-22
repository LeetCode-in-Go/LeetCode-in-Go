package s0215_kth_largest_element_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func TestFindKthLargest2(t *testing.T) {
	assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
