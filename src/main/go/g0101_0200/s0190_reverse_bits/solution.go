package s0190_reverse_bits

// #Easy #Top_Interview_Questions #Bit_Manipulation #Divide_and_Conquer
// #Algorithm_I_Day_14_Bit_Manipulation #Udemy_Bit_Manipulation #Top_Interview_150_Bit_Manipulation
// #2025_05_21_Time_0_ms_(100.00%)_Space_4.25_MB_(88.85%)

func reverseBits(num uint32) uint32 {
	var ret uint32
	// because there are 32 bits in total
	for i := 0; i < 32; i++ {
		ret = ret << 1
		// If the bit is 1 we OR it with 1, ie add 1
		if (num & 1) > 0 {
			ret = ret | 1
		}
		num = num >> 1
	}
	return ret
}
