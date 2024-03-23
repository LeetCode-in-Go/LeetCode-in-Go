package s0238_product_of_array_except_self

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Prefix_Sum
// #Data_Structure_II_Day_5_Array #Udemy_Arrays #Big_O_Time_O(n^2)_Space_O(n)
// #2024_03_23_Time_20_ms_(77.65%)_Space_7.8_MB_(37.12%)

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
