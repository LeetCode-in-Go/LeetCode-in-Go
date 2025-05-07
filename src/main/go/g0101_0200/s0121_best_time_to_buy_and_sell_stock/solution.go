package s0121_best_time_to_buy_and_sell_stock

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming
// #Data_Structure_I_Day_3_Array #Dynamic_Programming_I_Day_7 #Level_1_Day_5_Greedy #Udemy_Arrays
// #Top_Interview_150_Array/String #Big_O_Time_O(N)_Space_O(1)
// #2025_05_07_Time_0_ms_(100.00%)_Space_9.74_MB_(72.19%)

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices[1:] {
		if price > minPrice {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			minPrice = price
		}
	}
	return maxProfit
}
