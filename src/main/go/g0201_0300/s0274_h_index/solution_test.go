package s0274_h_index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHIndex(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	assert.Equal(t, 3, hIndex(citations))
}

func TestHIndex2(t *testing.T) {
	citations := []int{1, 3, 1}
	assert.Equal(t, 1, hIndex(citations))
}

func TestHIndex3(t *testing.T) {
	citations := []int{0}
	assert.Equal(t, 0, hIndex(citations))
}
