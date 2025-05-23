package s0219_contains_duplicate_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	assert.True(t, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func TestContainsNearbyDuplicate2(t *testing.T) {
	assert.True(t, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func TestContainsNearbyDuplicate3(t *testing.T) {
	assert.False(t, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
