package s0153_find_minimum_in_rotated_sorted_array

// #Medium #Top_100_Liked_Questions #Array #Binary_Search #Algorithm_II_Day_2_Binary_Search
// #Binary_Search_I_Day_12 #Udemy_Binary_Search #Top_Interview_150_Binary_Search
// #Big_O_Time_O(log_N)_Space_O(log_N) #2025_05_08_Time_0_ms_(100.00%)_Space_4.49_MB_(14.09%)

func findMin(nums []int) int {
	low, mid, high := 0, 0, len(nums)-1
	for low < high {
		mid = (low + high) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if nums[mid] < nums[low] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return nums[0]
}
