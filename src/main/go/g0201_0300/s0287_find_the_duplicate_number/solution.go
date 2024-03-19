package s0287_find_the_duplicate_number

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search #Two_Pointers
// #Bit_Manipulation #Binary_Search_II_Day_5 #Big_O_Time_O(n)_Space_O(n)
// #2024_03_18_Time_68_ms_(98.63%)_Space_8.3_MB_(69.20%)

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
