package s0763_partition_labels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	result := partitionLabels("ababcbacadefegdehijhklij")
	assert.Equal(t, []int{9, 7, 8}, result)
}

func TestPartitionLabels2(t *testing.T) {
	result := partitionLabels("eccbbbbdec")
	assert.Equal(t, []int{10}, result)
}
