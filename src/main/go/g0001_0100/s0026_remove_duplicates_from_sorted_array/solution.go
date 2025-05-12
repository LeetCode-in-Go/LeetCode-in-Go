package s0026_remove_duplicates_from_sorted_array

// #Easy #Top_Interview_Questions #Array #Two_Pointers #Udemy_Two_Pointers
// #Top_Interview_150_Array/String #2025_05_12_Time_0_ms_(100.00%)_Space_6.42_MB_(15.33%)

func removeDuplicates(nums []int) int {
	n := len(nums)
	i := 0
	j := 1
	if n <= 1 {
		return n
	}
	for j <= n-1 {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	return i + 1
}
