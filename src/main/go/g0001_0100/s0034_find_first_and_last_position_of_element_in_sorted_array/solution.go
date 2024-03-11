package s0034_find_first_and_last_position_of_element_in_sorted_array

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search
// #Algorithm_II_Day_1_Binary_Search #Binary_Search_I_Day_5 #Big_O_Time_O(log_n)_Space_O(1)
// #2024_03_11_Time_3_ms_(93.06%)_Space_4.5_MB_(100.00%)

func searchRange(nums []int, target int) []int {
	ans := make([]int, 2)
	ans[0] = helper(nums, target, false)
	ans[1] = helper(nums, target, true)
	return ans
}

func helper(nums []int, target int, equals bool) int {
	l, r := 0, len(nums)-1
	result := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			result = mid
		}
		if nums[mid] < target || (nums[mid] == target && equals) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}
