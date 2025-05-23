package s0212_word_search_ii

type Tree struct {
	children map[byte]*Tree
	end      string
}
