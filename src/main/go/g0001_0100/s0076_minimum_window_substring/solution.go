package s0076_minimum_window_substring

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Sliding_Window
// #Level_2_Day_14_Sliding_Window/Two_Pointer #Top_Interview_150_Sliding_Window
// #Big_O_Time_O(s.length())_Space_O(1) #2025_05_06_Time_0_ms_(100.00%)_Space_5.22_MB_(52.57%)

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
