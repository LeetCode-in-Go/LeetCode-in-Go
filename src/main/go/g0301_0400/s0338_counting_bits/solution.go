package s0338_counting_bits

// #Easy #Dynamic_Programming #Bit_Manipulation #LeetCode_75_Bit_Manipulation
// #Udemy_Bit_Manipulation #Big_O_Time_O(num)_Space_O(num)
// #2025_05_10_Time_0_ms_(100.00%)_Space_6.48_MB_(48.26%)

func countBits(num int) []int {
	result := make([]int, num+1)
	borderPos := 1
	incrPos := 1
	for i := 1; i < len(result); i++ {
		if incrPos == borderPos {
			result[i] = 1
			incrPos = 1
			borderPos = i
		} else {
			result[i] = 1 + result[incrPos]
			incrPos++
		}
	}
	return result
}
