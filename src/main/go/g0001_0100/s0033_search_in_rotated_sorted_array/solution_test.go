package s0033_search_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func TestSearch2(t *testing.T) {
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

func TestSearch3(t *testing.T) {
	assert.Equal(t, -1, search([]int{1}, 0))
}
