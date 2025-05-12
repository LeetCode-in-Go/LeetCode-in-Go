package s0027_remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	k := removeElement(nums, 3)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{2, 2, 2, 3}, nums)
}

func TestRemoveElement2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{0, 1, 4, 0, 3, 0, 4, 2}, nums)
}
