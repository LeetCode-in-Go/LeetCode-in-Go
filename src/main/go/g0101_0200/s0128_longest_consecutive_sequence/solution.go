package s0128_longest_consecutive_sequence

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table #Union_Find
// #Top_Interview_150_Hashmap #Big_O_Time_O(N_log_N)_Space_O(1)
// #2024_03_19_Time_54_ms_(99.84%)_Space_8.9_MB_(93.24%)
import "sort"


func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := -1 << 31 // Integer.MIN_VALUE in Go
	thsMax := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i]+1 {
			thsMax++
			continue
		}
		if nums[i+1] == nums[i] {
			continue
		}
		// Start of a new Sequence
		if thsMax > max {
			max = thsMax
		}
		thsMax = 1
	}
	if thsMax > max {
		max = thsMax
	}
	return max
}
