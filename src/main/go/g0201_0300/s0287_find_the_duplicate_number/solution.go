package s0287_find_the_duplicate_number

// #Medium #Top_100_Liked_Questions #Array #Binary_Search #Two_Pointers #Bit_Manipulation
// #Binary_Search_II_Day_5 #Big_O_Time_O(n)_Space_O(n)
// #2025_05_10_Time_1_ms_(90.38%)_Space_10.41_MB_(41.51%)

func findDuplicate(nums []int) int {
	arr := make([]int, len(nums)+1)
	for _, num := range nums {
		arr[num] += 1
		if arr[num] == 2 {
			return num
		}
	}
	return 0
}
