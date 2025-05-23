package s0209_minimum_size_subarray_sum

// #Medium #Array #Binary_Search #Prefix_Sum #Sliding_Window #Algorithm_II_Day_5_Sliding_Window
// #Binary_Search_II_Day_1 #Top_Interview_150_Sliding_Window
// #2025_05_23_Time_0_ms_(100.00%)_Space_9.32_MB_(64.50%)

func minSubArrayLen(target int, nums []int) int {
	i := 0
	j := 0
	sum := 0
	min := 1<<31 - 1
	for j < len(nums) {
		sum += nums[j]
		if sum >= target {
			for i <= j {
				if sum-nums[i] >= target {
					sum -= nums[i]
					i++
				} else {
					break
				}
			}
			if j-i+1 < min {
				min = j - i + 1
			}
		}
		j++
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}
