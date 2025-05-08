package s0198_house_robber

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming
// #LeetCode_75_DP/1D #Algorithm_I_Day_12_Dynamic_Programming #Dynamic_Programming_I_Day_3
// #Level_2_Day_12_Dynamic_Programming #Udemy_Dynamic_Programming #Top_Interview_150_1D_DP
// #Big_O_Time_O(n)_Space_O(n) #2025_05_08_Time_0_ms_(100.00%)_Space_3.88_MB_(99.40%)

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
