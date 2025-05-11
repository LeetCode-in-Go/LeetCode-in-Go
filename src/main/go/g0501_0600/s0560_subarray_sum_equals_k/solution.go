package s0560_subarray_sum_equals_k

// #Medium #Top_100_Liked_Questions #Array #Hash_Table #Prefix_Sum #Data_Structure_II_Day_5_Array
// #Big_O_Time_O(n)_Space_O(n) #2025_05_11_Time_11_ms_(93.74%)_Space_9.28_MB_(56.67%)

func subarraySum(nums []int, k int) int {
	tempSum := 0
	ret := 0
	sumCount := make(map[int]int)
	sumCount[0] = 1

	for _, num := range nums {
		tempSum += num
		if count, ok := sumCount[tempSum-k]; ok {
			ret += count
		}
		sumCount[tempSum]++
	}

	return ret
}
