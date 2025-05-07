package s0136_single_number

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Array #Bit_Manipulation
// #LeetCode_75_Bit_Manipulation #Data_Structure_II_Day_1_Array
// #Algorithm_I_Day_14_Bit_Manipulation #Udemy_Integers #Top_Interview_150_Bit_Manipulation
// #Big_O_Time_O(N)_Space_O(1) #2025_05_07_Time_0_ms_(100.00%)_Space_7.82_MB_(83.54%)

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
