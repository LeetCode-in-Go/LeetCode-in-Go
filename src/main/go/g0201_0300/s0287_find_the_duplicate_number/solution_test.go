package s0287_find_the_duplicate_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
}

func TestFindDuplicate2(t *testing.T) {
	assert.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
}

func TestFindDuplicate3(t *testing.T) {
	assert.Equal(t, 1, findDuplicate([]int{1, 1}))
}

func TestFindDuplicate4(t *testing.T) {
	assert.Equal(t, 1, findDuplicate([]int{1, 1, 2}))
}
