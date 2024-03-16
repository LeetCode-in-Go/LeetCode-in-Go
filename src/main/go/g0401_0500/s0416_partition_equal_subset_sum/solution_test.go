package s0416_partition_equal_subset_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanPartition(t *testing.T) {
	result := canPartition([]int{1, 5, 11, 5})
	assert.True(t, result)
}

func TestCanPartition2(t *testing.T) {
	result := canPartition([]int{1, 2, 3, 5})
	assert.False(t, result)
}
