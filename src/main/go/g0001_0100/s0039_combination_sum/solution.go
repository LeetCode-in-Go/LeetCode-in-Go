package s0039_combination_sum

// #Medium #Top_100_Liked_Questions #Array #Backtracking #Algorithm_II_Day_10_Recursion_Backtracking
// #Level_2_Day_20_Brute_Force/Backtracking #Udemy_Backtracking/Recursion
// #Top_Interview_150_Backtracking #Big_O_Time_O(2^n)_Space_O(n+2^n)
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.97_MB_(71.50%)

func combinationSum(coins []int, amount int) [][]int {
	var ans [][]int
	var subList []int
	combinationSumRec(len(coins), coins, amount, subList, &ans)
	return ans
}

func combinationSumRec(n int, coins []int, amount int, subList []int, ans *[][]int) {
	if amount == 0 || n == 0 {
		if amount == 0 {
			base := make([]int, len(subList))
			copy(base, subList)
			*ans = append(*ans, base)
		}
		return
	}
	if amount-coins[n-1] >= 0 {
		subList = append(subList, coins[n-1])
		combinationSumRec(n, coins, amount-coins[n-1], subList, ans)
		subList = subList[:len(subList)-1]
	}
	combinationSumRec(n-1, coins, amount, subList, ans)
}
