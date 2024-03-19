package s0136_single_number

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Array #Bit_Manipulation
// #Data_Structure_II_Day_1_Array #Algorithm_I_Day_14_Bit_Manipulation #Udemy_Integers
// #Big_O_Time_O(N)_Space_O(1) #2024_03_19_Time_8_ms_(95.49%)_Space_6.1_MB_(77.03%)

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
