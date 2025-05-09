package s0221_maximal_square

// #Medium #Array #Dynamic_Programming #Matrix #Dynamic_Programming_I_Day_16
// #Top_Interview_150_Multidimensional_DP #Big_O_Time_O(m*n)_Space_O(m*n)
// #2025_05_09_Time_1_ms_(93.36%)_Space_8.70_MB_(81.82%)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				// 1 + minimum from cell above, cell to the left, cell diagonal upper-left
				next := 1 + min(dp[i][j], min(dp[i+1][j], dp[i][j+1]))
				// keep track of the maximum value seen
				if next > max {
					max = next
				}
				dp[i+1][j+1] = next
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
