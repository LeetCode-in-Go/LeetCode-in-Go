package s0069_sqrtx

// #Easy #Top_Interview_Questions #Math #Binary_Search #Binary_Search_I_Day_4
// #Top_Interview_150_Math #2025_05_24_Time_0_ms_(100.00%)_Space_4.02_MB_(77.28%)

func mySqrt(x int) int {
	if x < 2 { // Handles x==0 and x==1 correctly
		return x
	}
	start := 1
	end := x / 2
	var sqrt int
	for start <= end {
		sqrt = start + (end-start)/2
		if sqrt == x/sqrt {
			return sqrt
		} else if sqrt > x/sqrt {
			end = sqrt - 1
		} else {
			start = sqrt + 1
		}
	}
	// sqrt will still contain the last value tried
	if sqrt > x/sqrt {
		return sqrt - 1
	}
	return sqrt
}
