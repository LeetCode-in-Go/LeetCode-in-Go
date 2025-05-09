package s0300_longest_increasing_subsequence

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Binary_Search
// #Algorithm_II_Day_16_Dynamic_Programming #Binary_Search_II_Day_3 #Dynamic_Programming_I_Day_18
// #Udemy_Dynamic_Programming #Top_Interview_150_1D_DP #Big_O_Time_O(n*log_n)_Space_O(n)
// #2025_05_10_Time_0_ms_(100.00%)_Space_5.28_MB_(69.25%)

import "math"

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	left, right := 1, 1
	for _, curr := range nums {
		start, end := left, right
		for start+1 < end {
			mid := start + (end-start)/2
			if dp[mid] > curr {
				end = mid
			} else {
				start = mid
			}
		}
		if dp[start] > curr {
			dp[start] = curr
		} else if curr > dp[start] && curr < dp[end] {
			dp[end] = curr
		} else if curr > dp[end] {
			dp[end+1] = curr
			right++
		}
	}
	return right
}
