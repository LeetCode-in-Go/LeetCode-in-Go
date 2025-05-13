package s0028_find_the_index_of_the_first_occurrence_in_a_string

// #Easy #Top_Interview_Questions #String #Two_Pointers #String_Matching
// #Programming_Skills_II_Day_1 #Top_Interview_150_Array/String
// #2025_05_13_Time_0_ms_(100.00%)_Space_3.83_MB_(94.91%)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	m := len(haystack)
	n := len(needle)
	for start := 0; start < m-n+1; start++ {
		if haystack[start:start+n] == needle {
			return start
		}
	}
	return -1
}
