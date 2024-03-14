package s0072_edit_distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
}

func TestMinDistance2(t *testing.T) {
	assert.Equal(t, 5, minDistance("intention", "execution"))
}
