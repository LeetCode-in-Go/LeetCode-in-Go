package s0030_substring_with_concatenation_of_all_words

// #Hard #String #Hash_Table #Sliding_Window #Top_Interview_150_Sliding_Window
// #2025_05_13_Time_5_ms_(91.71%)_Space_8.76_MB_(38.24%)

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	n1 := len(words[0])
	n2 := len(s)
	map1 := make(map[string]int)
	for _, ch := range words {
		map1[ch]++
	}
	for i := 0; i < n1; i++ {
		left := i
		j := i
		c := 0
		map2 := make(map[string]int)
		for j+n1 <= n2 {
			word1 := s[j : j+n1]
			j += n1
			if _, ok := map1[word1]; ok {
				map2[word1]++
				c++
				for map2[word1] > map1[word1] {
					word2 := s[left : left+n1]
					map2[word2]--
					left += n1
					c--
				}
				if c == len(words) {
					ans = append(ans, left)
				}
			} else {
				map2 = make(map[string]int)
				c = 0
				left = j
			}
		}
	}
	return ans
}
