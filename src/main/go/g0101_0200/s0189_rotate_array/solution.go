package s0189_rotate_array

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Math #Two_Pointers
// #Algorithm_I_Day_2_Two_Pointers #Udemy_Arrays #Big_O_Time_O(n)_Space_O(1)
// #2024_03_22_Time_20_ms_(92.61%)_Space_8.4_MB_(12.16%)

func reverse(nums []int, l, r int) {
	for l <= r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	t := n - (k % n)
	reverse(nums, 0, t-1)
	reverse(nums, t, n-1)
	reverse(nums, 0, n-1)
}
