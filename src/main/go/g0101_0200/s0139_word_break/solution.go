package s0139_word_break

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table
// #Dynamic_Programming #Trie #Memoization #Algorithm_II_Day_15_Dynamic_Programming
// #Dynamic_Programming_I_Day_9 #Udemy_Dynamic_Programming #Top_Interview_150_1D_DP
// #Big_O_Time_O(M+max*N)_Space_O(M+N+max) #2024_03_20_Time_0_ms_(100.00%)_Space_2.4_MB_(16.32%)

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}
	maxL := 0
	minL := len(wordDict[0])
	dict := map[string]bool{}
	for _, w := range wordDict {
		maxL = max(maxL, len(w))
		minL = min(minL, len(w))
		dict[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := minL; i <= len(s); i++ {
		for j := minL; j <= maxL && i-j >= 0; j++ {
			if dp[i-j] && dict[s[i-j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
