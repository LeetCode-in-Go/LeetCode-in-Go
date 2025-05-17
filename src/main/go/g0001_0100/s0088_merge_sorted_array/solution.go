package s0088_merge_sorted_array

// #Easy #Top_Interview_Questions #Array #Sorting #Two_Pointers #Data_Structure_I_Day_2_Array
// #Top_Interview_150_Array/String #2025_05_17_Time_0_ms_(100.00%)_Space_4.22_MB_(38.99%)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := len(nums1) - 1
	p2 := n - 1
	for p2 >= 0 {
		if i >= 0 && nums1[i] > nums2[p2] {
			nums1[j] = nums1[i]
			j--
			i--
		} else {
			nums1[j] = nums2[p2]
			j--
			p2--
		}
	}
}
