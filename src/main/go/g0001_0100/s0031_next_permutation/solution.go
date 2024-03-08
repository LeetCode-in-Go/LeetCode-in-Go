package s0031_next_permutation

// #Medium #Top_100_Liked_Questions #Array #Two_Pointers #Big_O_Time_O(n)_Space_O(1)
// #2024_03_08_Time_0_ms_(100.00%)_Space_2.4_MB_(7.18%)

func nextPermutation(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
