package s0005_longest_palindromic_substring

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming
// #Data_Structure_II_Day_9_String #Algorithm_II_Day_14_Dynamic_Programming
// #Dynamic_Programming_I_Day_17 #Udemy_Strings #Top_Interview_150_Multidimensional_DP
// #Big_O_Time_O(n)_Space_O(n) #2024_03_05_Time_0_ms_(100.00%)_Space_3.8_MB_(32.56%)

func longestPalindrome(s string) string {
	newStr := make([]byte, len(s)*2+1)
	newStr[0] = '#'
	for i := 0; i < len(s); i++ {
		newStr[2*i+1] = s[i]
		newStr[2*i+2] = '#'
	}
	dp := make([]int, len(newStr))
	friendCenter, friendRadius := 0, 0
	lpsCenter, lpsRadius := 0, 0

	for i := 0; i < len(newStr); i++ {
		if friendCenter+friendRadius > i {
			dp[i] = min(dp[2*friendCenter-i], (friendCenter+friendRadius)-i)
		} else {
			dp[i] = 1
		}
		for i+dp[i] < len(newStr) && i-dp[i] >= 0 && newStr[i+dp[i]] == newStr[i-dp[i]] {
			dp[i]++
		}
		if friendCenter+friendRadius < i+dp[i] {
			friendCenter, friendRadius = i, dp[i]
		}
		if lpsRadius < dp[i] {
			lpsCenter, lpsRadius = i, dp[i]
		}
	}
	return s[(lpsCenter-lpsRadius+1)/2 : (lpsCenter+lpsRadius-1)/2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
