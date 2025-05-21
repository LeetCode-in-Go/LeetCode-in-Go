package s0172_factorial_trailing_zeroes

// #Medium #Top_Interview_Questions #Math #Udemy_Integers #Top_Interview_150_Math
// #2025_05_21_Time_0_ms_(100.00%)_Space_3.90_MB_(100.00%)

func trailingZeroes(n int) int {
	base := 5
	count := 0
	for n >= base {
		count += n / base
		base = base * 5
	}
	return count
}
