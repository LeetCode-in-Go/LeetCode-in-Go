package s0188_best_time_to_buy_and_sell_stock_iv

// #Hard #Array #Dynamic_Programming #Top_Interview_150_Multidimensional_DP
// #2025_05_21_Time_0_ms_(100.00%)_Space_4.09_MB_(97.33%)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([]int, k+1)
	maxdp := make([]int, k+1)
	for i := 0; i <= k; i++ {
		maxdp[i] = -1 << 31
	}
	for i := 1; i <= n; i++ {
		maxdp[0] = max(maxdp[0], dp[0]-prices[i-1])
		for j := k; j >= 1; j-- {
			maxdp[j] = max(maxdp[j], dp[j]-prices[i-1])
			dp[j] = max(maxdp[j-1]+prices[i-1], dp[j])
		}
	}
	return dp[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
