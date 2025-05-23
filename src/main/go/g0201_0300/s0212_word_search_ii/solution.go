package s0212_word_search_ii

// #Hard #Top_Interview_Questions #Array #String #Matrix #Backtracking #Trie #Top_Interview_150_Trie
// #2025_05_23_Time_125_ms_(86.79%)_Space_6.63_MB_(97.17%)

var root *Tree

func (t *Tree) len() int {
	return len(t.children)
}

func (t *Tree) getChild(c byte) *Tree {
	return t.children[c]
}

func addWord(root *Tree, word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.children == nil {
			cur.children = make(map[byte]*Tree)
		}
		if cur.children[c] == nil {
			cur.children[c] = &Tree{}
		}
		cur = cur.children[c]
	}
	cur.end = word
}

func deleteWord(root *Tree, word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.children == nil {
			return
		}
		next := cur.children[c]
		if next == nil {
			return
		}
		if i == len(word)-1 {
			delete(cur.children, c)
		}
		cur = next
	}
}

func findWords(board [][]byte, words []string) []string {
	if len(board) < 1 || len(board[0]) < 1 {
		return []string{}
	}
	root = &Tree{}
	for _, word := range words {
		addWord(root, word)
	}
	collected := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root, &collected)
		}
	}
	return collected
}

func dfs(board [][]byte, i, j int, cur *Tree, collected *[]string) {
	c := board[i][j]
	if c == '-' {
		return
	}
	cur = cur.getChild(c)
	if cur == nil {
		return
	}
	if cur.end != "" {
		s := cur.end
		*collected = append(*collected, s)
		cur.end = ""
		if cur.len() == 0 {
			deleteWord(root, s)
		}
	}
	board[i][j] = '-'
	if i > 0 {
		dfs(board, i-1, j, cur, collected)
	}
	if i+1 < len(board) {
		dfs(board, i+1, j, cur, collected)
	}
	if j > 0 {
		dfs(board, i, j-1, cur, collected)
	}
	if j+1 < len(board[0]) {
		dfs(board, i, j+1, cur, collected)
	}
	board[i][j] = c
}

func decode(s string) string {
	return s
}
