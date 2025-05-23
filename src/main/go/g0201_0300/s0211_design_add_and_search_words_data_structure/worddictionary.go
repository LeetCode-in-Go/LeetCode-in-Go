package s0211_design_add_and_search_words_data_structure

// #Medium #String #Depth_First_Search #Design #Trie #Top_Interview_150_Trie
// #2025_05_23_Time_238_ms_(95.14%)_Space_73.80_MB_(25.72%)

type WordDictionary struct {
	root [26]*Node
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	n := len(word)
	if this.root[n] == nil {
		this.root[n] = &Node{}
	}
	node := this.root[n]
	for i := 0; i < n; i++ {
		c := word[i] - 'a'
		if node.kids[c] == nil {
			node.kids[c] = &Node{}
		}
		node = node.kids[c]
	}
	node.isTerminal = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this.root[len(word)]
	return node != nil && this.dfs(0, node, word)
}

func (this *WordDictionary) dfs(i int, node *Node, word string) bool {
	len := len(word)
	if i == len {
		return false
	}
	c := word[i]
	if c == '.' {
		for _, kid := range node.kids {
			if kid == nil {
				continue
			}
			if i == len-1 && kid.isTerminal || this.dfs(i+1, kid, word) {
				return true
			}
		}
		return false
	}
	kid := node.kids[c-'a']
	if kid == nil {
		return false
	}
	if i == len-1 {
		return kid.isTerminal
	}
	return this.dfs(i+1, kid, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
