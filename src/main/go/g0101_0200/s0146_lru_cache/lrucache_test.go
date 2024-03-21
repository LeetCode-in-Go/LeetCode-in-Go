package s0146_lru_cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	// cache is {1=1}
	cache.Put(1, 1)
	// cache is {1=1, 2=2}
	cache.Put(2, 2)
	// return 1
	assert.Equal(t, 1, cache.Get(1))
	// LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	cache.Put(3, 3)
	// returns -1 (not found)
	assert.Equal(t, -1, cache.Get(2))
	// LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	cache.Put(4, 4)
	// return -1 (not found)
	assert.Equal(t, -1, cache.Get(1))
	// return 3
	assert.Equal(t, 3, cache.Get(3))
	// return 4
	assert.Equal(t, 4, cache.Get(4))
}
