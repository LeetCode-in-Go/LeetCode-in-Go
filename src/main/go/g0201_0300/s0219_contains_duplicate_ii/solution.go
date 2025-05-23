package s0219_contains_duplicate_ii

// #Easy #Array #Hash_Table #Sliding_Window #Top_Interview_150_Hashmap
// #2025_05_23_Time_25_ms_(89.55%)_Space_18.84_MB_(9.71%)

func containsNearbyDuplicate(nums []int, k int) bool {
	duplicates := make(map[int]int, len(nums))
	i := len(nums) - 1
	for i >= 0 {
		if idx, seen := duplicates[nums[i]]; seen {
			if idx-i <= k {
				return true
			}
		}
		duplicates[nums[i]] = i
		i--
	}
	return false
}
