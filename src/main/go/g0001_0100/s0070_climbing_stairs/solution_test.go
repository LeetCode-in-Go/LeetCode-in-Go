package s0070_climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
}

func TestClimbStairs2(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
}
