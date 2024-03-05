package s0004_median_of_two_sorted_arrays

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search #Divide_and_Conquer
// #Big_O_Time_O(log(min(N,M)))_Space_O(1) #2024_03_05_Time_3_ms_(98.86%)_Space_4.8_MB_(98.47%)

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}

	var cut1, cut2, n1, n2 int
	n1 = len(nums1)
	n2 = len(nums2)
	low := 0
	high := n1

	for low <= high {
		cut1 = (low + high) / 2
		cut2 = ((n1 + n2 + 1) / 2) - cut1
		var l1, l2, r1, r2 int
		if cut1 == 0 {
			l1 = math.MinInt64
		} else {
			l1 = nums1[cut1-1]
		}
		if cut2 == 0 {
			l2 = math.MinInt64
		} else {
			l2 = nums2[cut2-1]
		}
		if cut1 == n1 {
			r1 = math.MaxInt64
		} else {
			r1 = nums1[cut1]
		}
		if cut2 == n2 {
			r2 = math.MaxInt64
		} else {
			r2 = nums2[cut2]
		}
		if l1 <= r2 && l2 <= r1 {
			if (n1+n2)%2 == 0 {
				return float64(math.Max(float64(l1),
					float64(l2))+math.Min(float64(r1), float64(r2))) / 2.0
			}
			return float64(math.Max(float64(l1), float64(l2)))
		} else if l1 > r2 {
			high = cut1 - 1
		} else {
			low = cut1 + 1
		}
	}
	return 0.0
}
