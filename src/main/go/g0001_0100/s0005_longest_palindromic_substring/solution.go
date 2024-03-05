package s0005_longest_palindromic_substring

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming
// #Data_Structure_II_Day_9_String #Algorithm_II_Day_14_Dynamic_Programming
// #Dynamic_Programming_I_Day_17 #Udemy_Strings #Big_O_Time_O(n)_Space_O(n)
// #2024_03_05_Time_0_ms_(100.00%)_Space_4.3_MB_(31.02%)

func longestPalindrome(s string) string {
	newStr := make([]rune, len(s)*2+1)
	newStr[0] = '#'
	for i, char := range s {
		newStr[2*i+1] = char
		newStr[2*i+2] = '#'
	}

	dp := make([]int, len(newStr))
	friendCenter := 0
	friendRadius := 0
	lpsCenter := 0
	lpsRadius := 0

	for i := 0; i < len(newStr); i++ {
		if friendCenter+friendRadius > i {
			dp[i] = min(dp[2*friendCenter-i], friendCenter+friendRadius-i)
		} else {
			dp[i] = 1
		}

		for i+dp[i] < len(newStr) && i-dp[i] >= 0 && newStr[i+dp[i]] == newStr[i-dp[i]] {
			dp[i]++
		}

		if friendCenter+friendRadius < i+dp[i] {
			friendCenter = i
			friendRadius = dp[i]
		}

		if lpsRadius < dp[i] {
			lpsCenter = i
			lpsRadius = dp[i]
		}
	}

	start := (lpsCenter - lpsRadius + 1) / 2
	end := (lpsCenter + lpsRadius - 1) / 2
	return s[start:end]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
