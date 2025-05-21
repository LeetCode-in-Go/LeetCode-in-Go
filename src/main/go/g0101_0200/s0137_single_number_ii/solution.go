package s0137_single_number_ii

// #Medium #Array #Bit_Manipulation #Top_Interview_150_Bit_Manipulation
// #2025_05_21_Time_0_ms_(100.00%)_Space_5.16_MB_(75.27%)

func singleNumber(nums []int) int {
	ones := 0
	twos := 0
	for _, num := range nums {
		ones = (ones ^ num) & (^twos)
		twos = (twos ^ num) & (^ones)
	}
	return ones
}
