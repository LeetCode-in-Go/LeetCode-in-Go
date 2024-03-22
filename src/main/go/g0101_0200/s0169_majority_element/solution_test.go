package s0169_majority_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
}

func TestMajorityElement2(t *testing.T) {
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
