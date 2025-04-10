package s0053_maximum_subarray

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming
// #Divide_and_Conquer #Data_Structure_I_Day_1_Array #Dynamic_Programming_I_Day_5
// #Udemy_Famous_Algorithm #Top_Interview_150_Kadane's_Algorithm #Big_O_Time_O(n)_Space_O(1)
// #2024_03_13_Time_79_ms_(87.63%)_Space_7.8_MB_(92.56%)

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return max
}
