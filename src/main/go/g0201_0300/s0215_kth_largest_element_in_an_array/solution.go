package s0215_kth_largest_element_in_an_array

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Sorting #Heap_Priority_Queue
// #Divide_and_Conquer #Quickselect #LeetCode_75_Heap/Priority_Queue
// #Data_Structure_II_Day_20_Heap_Priority_Queue #Top_Interview_150_Heap
// #Big_O_Time_O(n*log(n))_Space_O(log(n)) #2025_05_09_Time_17_ms_(83.49%)_Space_10.00_MB_(66.67%)

import "sort"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	return nums[n-k]
}
