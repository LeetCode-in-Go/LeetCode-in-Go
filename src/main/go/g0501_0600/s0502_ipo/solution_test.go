package s0502_ipo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaximizedCapital(t *testing.T) {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	assert.Equal(t, 4, findMaximizedCapital(k, w, profits, capital))
}

func TestFindMaximizedCapital2(t *testing.T) {
	k := 3
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 2}
	assert.Equal(t, 6, findMaximizedCapital(k, w, profits, capital))
}

func TestFindMaximizedCapital3(t *testing.T) {
	k := 1
	w := 2
	profits := []int{1, 2, 3}
	capital := []int{1, 1, 2}
	assert.Equal(t, 5, findMaximizedCapital(k, w, profits, capital))
}

func TestFindMaximizedCapital4(t *testing.T) {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 9, 10}
	assert.Equal(t, 1, findMaximizedCapital(k, w, profits, capital))
}
