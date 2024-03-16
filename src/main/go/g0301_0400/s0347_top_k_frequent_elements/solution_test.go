package s0347_top_k_frequent_elements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	result := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	expected := []int{1, 2}
	assert.Equal(t, expected, result)
}

func TestTopKFrequent2(t *testing.T) {
	result := topKFrequent([]int{1}, 1)
	expected := []int{1}
	assert.Equal(t, expected, result)
}
