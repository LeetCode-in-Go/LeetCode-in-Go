package s0918_maximum_sum_circular_subarray

// #Medium #Array #Dynamic_Programming #Divide_and_Conquer #Queue #Monotonic_Queue
// #Dynamic_Programming_I_Day_5 #Top_Interview_150_Kadane's_Algorithm
// #2025_05_24_Time_0_ms_(100.00%)_Space_9.13_MB_(33.13%)

func maxSubarraySumCircular(nums []int) int {
	total := 0
	maxSum := nums[0]
	curMax := 0
	minSum := nums[0]
	curMin := 0
	for _, num := range nums {
		total += num
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)
	}
	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
