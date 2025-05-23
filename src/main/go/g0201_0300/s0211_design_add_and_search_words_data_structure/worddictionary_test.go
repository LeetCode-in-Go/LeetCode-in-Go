package s0211_design_add_and_search_words_data_structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	input := []string{"bad", "dad", "mad"}
	dictionary := Constructor()
	for _, s := range input {
		dictionary.AddWord(s)
	}
	assert.False(t, dictionary.Search("pad"))
	assert.True(t, dictionary.Search("bad"))
	assert.True(t, dictionary.Search(".ad"))
	assert.True(t, dictionary.Search("b.."))
}
