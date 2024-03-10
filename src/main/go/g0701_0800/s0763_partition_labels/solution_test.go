package s0763_partition_labels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionLabels1(t *testing.T) {
	assert := assert.New(t)
	result := partitionLabels("ababcbacadefegdehijhklij")
	assert.Equal([]int{9, 7, 8}, result)
}

func TestPartitionLabels2(t *testing.T) {
	assert := assert.New(t)
	result := partitionLabels("eccbbbbdec")
	assert.Equal([]int{10}, result)
}
