package s0239_sliding_window_maximum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	assert.Equal(
		t,
		[]int{3, 3, 5, 5, 6, 7},
		maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)
}

func TestMaxSlidingWindow2(t *testing.T) {
	assert.Equal(
		t,
		[]int{1},
		maxSlidingWindow([]int{1}, 1),
	)
}
