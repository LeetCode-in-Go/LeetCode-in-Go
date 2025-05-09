package s0238_product_of_array_except_self

// #Medium #Top_100_Liked_Questions #Array #Prefix_Sum #LeetCode_75_Array/String
// #Data_Structure_II_Day_5_Array #Udemy_Arrays #Top_Interview_150_Array/String
// #Big_O_Time_O(n^2)_Space_O(n) #2025_05_09_Time_0_ms_(100.00%)_Space_9.70_MB_(86.27%)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	product := 1
	ans := make([]int, n)

	// Calculate the product of all elements in the array
	for _, num := range nums {
		product *= num
	}

	// Calculate the result array
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			ans[i] = product / nums[i]
		} else {
			p := 1
			for j := 0; j < n; j++ {
				if j != i {
					p *= nums[j]
				}
			}
			ans[i] = p
		}
	}
	return ans
}
