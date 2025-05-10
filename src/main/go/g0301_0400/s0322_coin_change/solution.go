package s0322_coin_change

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Breadth_First_Search
// #Algorithm_II_Day_18_Dynamic_Programming #Dynamic_Programming_I_Day_20
// #Level_2_Day_12_Dynamic_Programming #Top_Interview_150_1D_DP #Big_O_Time_O(m*n)_Space_O(amount)
// #2025_05_10_Time_10_ms_(79.59%)_Space_8.61_MB_(51.21%)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i < c {
				continue
			}

			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
