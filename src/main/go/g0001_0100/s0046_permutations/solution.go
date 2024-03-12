package s0046_permutations

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Backtracking
// #Algorithm_I_Day_11_Recursion_Backtracking #Level_2_Day_20_Brute_Force/Backtracking
// #Udemy_Backtracking/Recursion #Big_O_Time_O(n*n!)_Space_O(n+n!)
// #2024_03_12_Time_0_ms_(100.00%)_Space_2.7_MB_(51.79%)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	finalResult := [][]int{}
	permuteRecur(nums, &finalResult, []int{}, make([]bool, len(nums)))
	return finalResult
}

func permuteRecur(nums []int, finalResult *[][]int, currResult []int, used []bool) {
	if len(currResult) == len(nums) {
		temp := make([]int, len(currResult))
		copy(temp, currResult)
		*finalResult = append(*finalResult, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		currResult = append(currResult, nums[i])
		used[i] = true
		permuteRecur(nums, finalResult, currResult, used)
		used[i] = false
		currResult = currResult[:len(currResult)-1]
	}
}
