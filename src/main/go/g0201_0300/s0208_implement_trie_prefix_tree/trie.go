package s0208_implement_trie_prefix_tree

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Design #Trie
// #LeetCode_75_Trie #Level_2_Day_16_Design #Udemy_Trie_and_Heap #Top_Interview_150_Trie
// #Big_O_Time_O(word.length())_or_O(prefix.length())_Space_O(N)
// #2025_05_09_Time_11_ms_(98.49%)_Space_16.18_MB_(42.49%)

import "strings"

type node struct {
	val   string
	left  *node
	right *node
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	t.root = insert(t.root, word)
}

func (t *Trie) Search(word string) bool {
	return search(t.root, word)
}

func (t *Trie) StartsWith(prefix string) bool {
	return startsWith(t.root, prefix)
}

func startsWith(root *node, prefix string) bool {
	if root == nil {
		return false
	}
	if strings.Index(root.val, prefix) == 0 {
		return true
	}
	if prefix < root.val {
		return startsWith(root.left, prefix)
	}
	return startsWith(root.right, prefix)
}

func search(root *node, word string) bool {
	if root == nil {
		return false
	}
	if word == root.val {
		return true
	}
	if word < root.val {
		return search(root.left, word)
	}
	return search(root.right, word)
}

func insert(root *node, word string) *node {
	if root == nil {
		return &node{word, nil, nil}
	}
	if word == root.val {
		return root
	}
	if word < root.val {
		root.left = insert(root.left, word)
	} else {
		root.right = insert(root.right, word)
	}
	return root
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
