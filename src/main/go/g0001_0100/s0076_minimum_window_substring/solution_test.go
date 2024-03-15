package s0076_minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
}

func TestMinWindow2(t *testing.T) {
	assert.Equal(t, "a", minWindow("a", "a"))
}

func TestMinWindow3(t *testing.T) {
	assert.Equal(t, "", minWindow("a", "aa"))
}
