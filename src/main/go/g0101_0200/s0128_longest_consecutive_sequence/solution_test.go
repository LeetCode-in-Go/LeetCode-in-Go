package s0128_longest_consecutive_sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func TestLongestConsecutive2(t *testing.T) {
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
