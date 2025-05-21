package s0162_find_peak_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1}))
}

func TestFindPeakElement2(t *testing.T) {
	assert.Equal(t, 5, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
