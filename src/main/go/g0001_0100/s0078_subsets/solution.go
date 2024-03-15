package s0078_subsets

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Bit_Manipulation #Backtracking
// #Algorithm_II_Day_9_Recursion_Backtracking #Udemy_Backtracking/Recursion
// #Big_O_Time_O(2^n)_Space_O(n*2^n) #2024_03_15_Time_1_ms_(80.53%)_Space_2.4_MB_(60.25%)

func subsets(nums []int) [][]int {
	var res [][]int
	solve(nums, []int{}, &res, 0)
	return res
}

// solve is a recursive helper function to generate subsets.
func solve(nums []int, temp []int, res *[][]int, start int) {
	*res = append(*res, append([]int(nil), temp...))
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		solve(nums, temp, res, i+1)
		temp = temp[:len(temp)-1]
	}
}
