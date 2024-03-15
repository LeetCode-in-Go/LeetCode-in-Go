package s0096_unique_binary_search_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTrees(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
}

func TestNumTrees2(t *testing.T) {
	assert.Equal(t, 1, numTrees(1))
}
