package s0416_partition_equal_subset_sum

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Level_2_Day_13_Dynamic_Programming
// #Big_O_Time_O(n*sums)_Space_O(n*sums) #2025_05_10_Time_4_ms_(97.02%)_Space_4.42_MB_(72.07%)

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum, sum)
	dp[0] = true
	for _, n := range nums {
		if n <= sum {
			if dp[sum-n] == true {
				return true
			}
			for j := sum - n - 1; j >= 0; j-- {
				if dp[j] == true {
					dp[j+n] = true
				}
			}
		}
	}
	return false
}
