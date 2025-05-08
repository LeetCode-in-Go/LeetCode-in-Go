package s0169_majority_element

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table #Sorting #Counting
// #Divide_and_Conquer #Data_Structure_II_Day_1_Array #Udemy_Famous_Algorithm
// #Top_Interview_150_Array/String #Big_O_Time_O(n)_Space_O(1)
// #2025_05_08_Time_0_ms_(100.00%)_Space_8.53_MB_(56.86%)

func majorityElement(nums []int) int {
	count := 1
	majority := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			count++
		} else {
			if count > 1 {
				count--
			} else {
				majority = nums[i]
			}
		}
	}
	count = 0
	for _, num := range nums {
		if num == majority {
			count++
		}
	}
	if count >= (len(nums)/2)+1 {
		return majority
	} else {
		return -1
	}
}
