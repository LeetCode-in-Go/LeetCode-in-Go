package s0014_longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func TestLongestCommonPrefix2(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
