package s0027_remove_element

// #Easy #Array #Two_Pointers #Top_Interview_150_Array/String
// #2025_05_12_Time_0_ms_(100.00%)_Space_4.12_MB_(19.61%)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	len := len(nums)
	j := len - 1
	occurTimes := 0
	for i := 0; i < len; i++ {
		if nums[i] == val {
			occurTimes++
			if j == i {
				return len - occurTimes
			}
			for nums[j] == val {
				j--
				occurTimes++
				if j == i {
					return len - occurTimes
				}
			}
			nums[i] = nums[j]
			j--
		}
		if i == j {
			return len - occurTimes
		}
	}
	return len - occurTimes
}
