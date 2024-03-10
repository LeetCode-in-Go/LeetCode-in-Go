package s0763_partition_labels

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels1(t *testing.T) {
	result := partitionLabels("ababcbacadefegdehijhklij")
	assert.Equal(t, []int{9, 7, 8}, result)
}

func TestPartitionLabels2(t *testing.T) {
	result := partitionLabels("eccbbbbdec")
	assert.Equal(t, []int{10}, result)
}
