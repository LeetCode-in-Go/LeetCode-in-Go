package s0050_powx_n

// #Medium #Top_Interview_Questions #Math #Recursion #Udemy_Integers #Top_Interview_150_Math
// #2025_05_24_Time_0_ms_(100.00%)_Space_3.88_MB_(100.00%)

func myPow(x float64, n int) float64 {
	nn := int64(n)
	res := 1.0
	if n < 0 {
		nn = -nn
	}
	for nn > 0 {
		if nn%2 == 1 {
			nn--
			res *= x
		} else {
			x *= x
			nn /= 2
		}
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
