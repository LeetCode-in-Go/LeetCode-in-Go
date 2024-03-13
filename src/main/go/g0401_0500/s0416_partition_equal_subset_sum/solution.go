package s0416_partition_equal_subset_sum

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Level_2_Day_13_Dynamic_Programming
// #Big_O_Time_O(n*sums)_Space_O(n*sums) #2024_03_12_Time_13_ms_(91.83%)_Space_2.5_MB_(98.44%)

func canPartition(nums []int) bool {
	sums := 0
	for _, num := range nums {
		sums += num
	}
	if sums%2 == 1 {
		return false
	}
	sums /= 2
	dp := make([]bool, sums+1)
	dp[0] = true
	for _, num := range nums {
		for sum := sums; sum >= num; sum-- {
			dp[sum] = dp[sum] || dp[sum-num]
		}
	}
	return dp[sums]
}
