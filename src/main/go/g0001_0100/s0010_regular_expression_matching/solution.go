package s0010_regular_expression_matching

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming #Recursion
// #Udemy_Dynamic_Programming #Big_O_Time_O(m*n)_Space_O(m*n)
// #2024_03_07_Time_0_ms_(100.00%)_Space_2.3_MB_(36.02%)

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

			}
		}
	}
	return dp[m][n]
}
