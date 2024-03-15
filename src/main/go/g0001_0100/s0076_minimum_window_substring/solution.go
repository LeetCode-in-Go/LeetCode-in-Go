package s0076_minimum_window_substring

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Sliding_Window
// #Level_2_Day_14_Sliding_Window/Two_Pointer #Big_O_Time_O(s.length())_Space_O(1)
// #2024_03_15_Time_0_ms_(100.00%)_Space_3_MB_(66.67%)

import "math"

func minWindow(s string, t string) string {
	mapT := make([]int, 128)
	for i := 0; i < len(t); i++ {
		mapT[t[i]-'A']++
	}
	count := len(t)
	begin, end, d, head := 0, 0, math.MaxInt32, 0

	for end < len(s) {
		if mapT[s[end]-'A']--; mapT[s[end]-'A'] >= 0 {
			count--
		}
		end++

		for count == 0 {
			if end-begin < d {
				d = end - begin
				head = begin
			}
			if mapT[s[begin]-'A']++; mapT[s[begin]-'A'] > 0 {
				count++
			}
			begin++
		}
	}

	if d == math.MaxInt32 {
		return ""
	}
	return s[head : head+d]
}
