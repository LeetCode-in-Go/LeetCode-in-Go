package s0062_unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
}

func TestUniquePaths2(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
}
