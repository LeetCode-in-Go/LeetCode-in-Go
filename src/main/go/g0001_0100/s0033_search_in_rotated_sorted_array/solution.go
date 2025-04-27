package s0033_search_in_rotated_sorted_array

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search
// #Algorithm_II_Day_1_Binary_Search #Binary_Search_I_Day_11 #Level_2_Day_8_Binary_Search
// #Udemy_Binary_Search #Top_Interview_150_Binary_Search #Big_O_Time_O(log_n)_Space_O(1)
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.51_MB_(4.93%)

func search(nums []int, target int) int {
	var mid, lo, hi int
	hi = len(nums) - 1
	for lo <= hi {
		mid = ((hi - lo) >> 1) + lo
		if target == nums[mid] {
			return mid
		}
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target <= nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] <= target && target <= nums[hi] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
