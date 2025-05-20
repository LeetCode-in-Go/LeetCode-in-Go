package s0127_word_ladder

// #Hard #Top_Interview_Questions #String #Hash_Table #Breadth_First_Search
// #Graph_Theory_I_Day_12_Breadth_First_Search #Top_Interview_150_Graph_BFS
// #2025_05_20_Time_17_ms_(97.47%)_Space_8.67_MB_(78.75%)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	beginSet := make(map[string]bool)
	endSet := make(map[string]bool)
	visited := make(map[string]bool)
	beginSet[beginWord] = true
	endSet[endWord] = true
	lenght := 1
	strLen := len(beginWord)
	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		tempSet := make(map[string]bool)
		for s := range beginSet {
			chars := []byte(s)
			for i := 0; i < strLen; i++ {
				old := chars[i]
				for j := byte('a'); j <= 'z'; j++ {
					chars[i] = j
					temp := string(chars)
					if endSet[temp] {
						return lenght + 1
					}
					if !visited[temp] && wordSet[temp] {
						tempSet[temp] = true
						visited[temp] = true
					}
				}
				chars[i] = old
			}
		}
		beginSet = tempSet
		lenght++
	}
	return 0
}
