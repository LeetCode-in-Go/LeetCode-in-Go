package s0120_triangle

// #Medium #Array #Dynamic_Programming #Algorithm_I_Day_12_Dynamic_Programming
// #Dynamic_Programming_I_Day_13 #Udemy_Dynamic_Programming #Top_Interview_150_Multidimensional_DP
// #2025_05_18_Time_0_ms_(100.00%)_Space_5.48_MB_(16.19%)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[len(triangle)-1]))
		for j := range dp[i] {
			dp[i][j] = -10001
		}
	}
	return dfs(triangle, dp, 0, 0)
}

func dfs(triangle [][]int, dp [][]int, row, col int) int {
	if row >= len(triangle) {
		return 0
	}
	if dp[row][col] != -10001 {
		return dp[row][col]
	}
	sum := triangle[row][col] + min(dfs(triangle, dp, row+1, col), dfs(triangle, dp, row+1, col+1))
	dp[row][col] = sum
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
