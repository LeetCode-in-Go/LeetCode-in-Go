package s0494_target_sum

// #Medium #Array #Dynamic_Programming #Backtracking #Big_O_Time_O(n*(sum+s))_Space_O(n*(sum+s))
// #2025_05_11_Time_3_ms_(84.35%)_Space_8.52_MB_(36.96%)

import "math"

func findTargetSumWays(nums []int, s int) int {
	sum := 0
	s = int(math.Abs(float64(s)))
	for _, num := range nums {
		sum += num
	}
	if s > sum || (sum+s)%2 != 0 {
		return 0
	}
	dp := make([][]int, (sum+s)/2+1)
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			dp[0][i+1] = dp[0][i] * 2
		} else {
			dp[0][i+1] = dp[0][i]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(nums); j++ {
			dp[i][j+1] += dp[i][j]
			if nums[j] <= i {
				dp[i][j+1] += dp[i-nums[j]][j]
			}
		}
	}
	return dp[(sum+s)/2][len(nums)]
}
