package s0191_number_of_1_bits

// #Easy #Top_Interview_Questions #Bit_Manipulation #Algorithm_I_Day_13_Bit_Manipulation
// #Programming_Skills_I_Day_2_Operator #Udemy_Bit_Manipulation #Top_Interview_150_Bit_Manipulation
// #2025_05_21_Time_0_ms_(100.00%)_Space_3.96_MB_(73.70%)

func hammingWeight(num int) int {
	sum := 0
	for num > 0 {
		sum += int(num & 1)
		num >>= 1
	}
	return sum
}
