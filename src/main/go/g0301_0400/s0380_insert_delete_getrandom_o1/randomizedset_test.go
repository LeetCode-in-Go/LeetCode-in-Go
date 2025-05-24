package s0380_insert_delete_getrandom_o1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(1))
	assert.False(t, randomizedSet.Remove(2))
	assert.True(t, randomizedSet.Insert(2))
	randomizedSet.GetRandom()
	assert.True(t, randomizedSet.Remove(1))
	assert.False(t, randomizedSet.Insert(2))
	assert.Equal(t, 2, randomizedSet.GetRandom())
}

func TestRandomizedSet2(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(0))
	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Remove(0))
	assert.True(t, randomizedSet.Insert(2))
	assert.True(t, randomizedSet.Remove(1))
	assert.Equal(t, 2, randomizedSet.GetRandom())
}

func TestRandomizedSet3(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(3))
	assert.False(t, randomizedSet.Insert(3))
	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Remove(3))
	assert.True(t, randomizedSet.Insert(0))
	assert.True(t, randomizedSet.Remove(0))
}
