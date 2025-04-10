package s0072_edit_distance

// #Medium #Top_100_Liked_Questions #String #Dynamic_Programming
// #Algorithm_II_Day_18_Dynamic_Programming #Dynamic_Programming_I_Day_19
// #Udemy_Dynamic_Programming #Top_Interview_150_Multidimensional_DP #Big_O_Time_O(n^2)_Space_O(n2)
// #2024_03_14_Time_0_ms_(100.00%)_Space_2.6_MB_(97.61%)

func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	if n2 > n1 {
		return minDistance(word2, word1)
	}
	dp := make([]int, n2+1)
	for j := 0; j <= n2; j++ {
		dp[j] = j
	}
	for i := 1; i <= n1; i++ {
		prev := dp[0]
		dp[0] = i
		for j := 1; j <= n2; j++ {
			temp := dp[j]
			if word1[i-1] != word2[j-1] {
				dp[j] = 1 + min(prev, min(dp[j], dp[j-1]))
			} else {
				dp[j] = prev
			}
			prev = temp
		}
	}
	return dp[n2]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
