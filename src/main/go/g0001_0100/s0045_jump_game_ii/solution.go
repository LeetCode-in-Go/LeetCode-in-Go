package s0045_jump_game_ii

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Greedy
// #Algorithm_II_Day_13_Dynamic_Programming #Dynamic_Programming_I_Day_4
// #Top_Interview_150_Array/String #Big_O_Time_O(n)_Space_O(1)
// #2024_03_12_Time_6_ms_(96.98%)_Space_6.3_MB_(34.53%)

func jump(nums []int) int {
	length := 0
	maxLength := 0
	minJump := 0
	for i := 0; i < len(nums)-1; i++ {
		length--
		maxLength--
		maxLength = max(maxLength, nums[i])
		if length <= 0 {
			length = maxLength
			minJump++
		}
		if length >= len(nums)-i-1 {
			return minJump
		}
	}
	return minJump
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
