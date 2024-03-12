package s0042_trapping_rain_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func TestTrap2(t *testing.T) {
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}
