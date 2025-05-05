package s0055_jump_game

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming #Greedy
// #Algorithm_II_Day_12_Dynamic_Programming #Dynamic_Programming_I_Day_4 #Udemy_Arrays
// #Top_Interview_150_Array/String #Big_O_Time_O(n)_Space_O(1)
// #2025_05_05_Time_0_ms_(100.00%)_Space_8.96_MB_(66.09%)

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	jumpsNeeded := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < jumpsNeeded {
			jumpsNeeded += 1
		} else {
			jumpsNeeded = 1
		}
	}

	return jumpsNeeded == 1
}
