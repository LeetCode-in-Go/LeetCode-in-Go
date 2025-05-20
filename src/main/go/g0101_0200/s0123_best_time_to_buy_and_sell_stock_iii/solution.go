package s0123_best_time_to_buy_and_sell_stock_iii

// #Hard #Array #Dynamic_Programming #Top_Interview_150_Multidimensional_DP
// #2025_05_20_Time_0_ms_(100.00%)_Space_12.50_MB_(32.12%)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	fb := -1 << 31
	sb := -1 << 31
	fs := 0
	ss := 0
	for _, price := range prices {
		fb = max(fb, -price)
		fs = max(fs, fb+price)
		sb = max(sb, fs-price)
		ss = max(ss, sb+price)
	}
	return ss
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
