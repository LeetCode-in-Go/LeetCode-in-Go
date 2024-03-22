package s0198_house_robber

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming
// #Algorithm_I_Day_12_Dynamic_Programming #Dynamic_Programming_I_Day_3
// #Level_2_Day_12_Dynamic_Programming #Udemy_Dynamic_Programming #Big_O_Time_O(n)_Space_O(n)
// #2024_03_22_Time_1_ms_(78.69%)_Space_2.3_MB_(7.60%)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	profit := make([]int, len(nums))
	profit[0] = nums[0]
	profit[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		profit[i] = max(profit[i-1], nums[i]+profit[i-2])
	}
	return profit[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
