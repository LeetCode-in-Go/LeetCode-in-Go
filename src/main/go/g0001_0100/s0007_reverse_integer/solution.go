package s0007_reverse_integer

// #Medium #Top_Interview_Questions #Math #Udemy_Integers
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.03_MB_(24.37%)

import "math"

func reverse(x int) int {
	var r int
	for ok := true; ok; ok = x != 0 {
		if r > math.MaxInt32/10 || r < math.MinInt32/10 {
			return 0
		}
		r = (x % 10) + (r * 10)
		x /= 10
	}
	return r
}
