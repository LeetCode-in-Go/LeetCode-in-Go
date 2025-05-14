package s0080_remove_duplicates_from_sorted_array_ii

// #Medium #Array #Two_Pointers #Udemy_Arrays #Top_Interview_150_Array/String
// #2025_05_14_Time_6_ms_(77.80%)_Space_7.04_MB_(83.23%)

func removeDuplicates(nums []int) int {
	i := 0
	k := 0
	count := 0
	for i < len(nums)-1 {
		count++
		if count <= 2 {
			nums[k] = nums[i]
			k++
		}
		if nums[i] != nums[i+1] {
			count = 0
			i++
			continue
		}
		i++
	}
	count++
	if count <= 2 {
		nums[k] = nums[i]
		k++
	}
	return k
}
