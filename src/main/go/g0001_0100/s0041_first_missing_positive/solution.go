package s0041_first_missing_positive

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table #Udemy_Arrays
// #Big_O_Time_O(n)_Space_O(n) #2025_05_03_Time_0_ms_(100.00%)_Space_9.49_MB_(95.52%)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		val := abs(nums[i])
		if val > 0 && val <= len(nums) {
			if nums[val-1] > 0 {
				nums[val-1] = nums[val-1] * -1
				continue
			}
			if nums[val-1] == 0 {
				nums[val-1] = -1 * (len(nums) + 1)
			}
		}
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] >= 0 {
			return i
		}
	}
	return len(nums) + 1
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * -1
}
