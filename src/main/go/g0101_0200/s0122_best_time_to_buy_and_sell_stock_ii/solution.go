package s0122_best_time_to_buy_and_sell_stock_ii

// #Medium #Top_Interview_Questions #Array #Dynamic_Programming #Greedy #Dynamic_Programming_I_Day_7
// #Udemy_Arrays #Top_Interview_150_Array/String
// #2025_05_18_Time_0_ms_(100.00%)_Space_5.11_MB_(46.22%)

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
