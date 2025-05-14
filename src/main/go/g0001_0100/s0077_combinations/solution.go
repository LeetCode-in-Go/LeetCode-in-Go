package s0077_combinations

// #Medium #Backtracking #Algorithm_I_Day_11_Recursion_Backtracking #Top_Interview_150_Backtracking
// #2025_05_14_Time_36_ms_(77.01%)_Space_54.93_MB_(82.35%)

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	if n > 20 || k < 1 || k > n {
		return ans
	}
	backtrack(&ans, n, k, 1, make([]int, 0))
	return ans
}

func backtrack(ans *[][]int, n, k, s int, stack []int) {
	if k == 0 {
		*ans = append(*ans, append([]int{}, stack...))
		return
	}
	for i := s; i <= (n-k)+1; i++ {
		stack = append(stack, i)
		backtrack(ans, n, k-1, i+1, stack)
		stack = stack[:len(stack)-1]
	}
}
