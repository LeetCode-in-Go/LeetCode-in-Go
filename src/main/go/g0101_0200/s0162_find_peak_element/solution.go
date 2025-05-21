package s0162_find_peak_element

// #Medium #Top_Interview_Questions #Array #Binary_Search #LeetCode_75_Binary_Search
// #Algorithm_II_Day_2_Binary_Search #Binary_Search_II_Day_12 #Top_Interview_150_Binary_Search
// #2025_05_21_Time_0_ms_(100.00%)_Space_4.72_MB_(1.24%)

func findPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		// This is done because start and end might be big numbers, so it might exceed the
		// integer limit.
		mid := start + ((end - start) / 2)
		if nums[mid+1] > nums[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
