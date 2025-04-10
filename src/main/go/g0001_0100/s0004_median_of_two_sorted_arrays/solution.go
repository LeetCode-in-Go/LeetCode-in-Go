package s0004_median_of_two_sorted_arrays

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search #Divide_and_Conquer
// #Top_Interview_150_Binary_Search #Big_O_Time_O(log(min(N,M)))_Space_O(1)
// #AI_can_be_used_to_solve_the_task #2024_03_05_Time_9_ms_(72.04%)_Space_4.8_MB_(67.69%)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := getIntValue(nums1, partitionX-1)
		minRightX := getIntValue(nums1, partitionX)

		maxLeftY := getIntValue(nums2, partitionY-1)
		minRightY := getIntValue(nums2, partitionY)

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}

func getIntValue(nums []int, index int) int {
	if index < 0 {
		return -1000000
	}
	if index >= len(nums) {
		return 1000000
	}
	return nums[index]
}
