package s0209_minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func TestMinSubArrayLen2(t *testing.T) {
	assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
}

func TestMinSubArrayLen3(t *testing.T) {
	assert.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
