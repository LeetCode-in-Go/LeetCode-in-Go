package s0283_move_zeroes

// #Easy #Top_100_Liked_Questions #Array #Two_Pointers #LeetCode_75_Two_Pointers
// #Algorithm_I_Day_3_Two_Pointers #Programming_Skills_I_Day_6_Array #Udemy_Arrays
// #Big_O_Time_O(n)_Space_O(1) #2025_05_10_Time_0_ms_(100.00%)_Space_9.28_MB_(18.25%)

func moveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
