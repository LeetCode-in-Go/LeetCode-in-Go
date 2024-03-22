package s0208_implement_trie_prefix_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	// return True
	assert.True(t, trie.Search("apple"))
	// return False
	assert.False(t, trie.Search("app"))
	// return True
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	// return True
	assert.True(t, trie.Search("app"))
}
