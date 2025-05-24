package s0392_is_subsequence

// #Easy #String #Dynamic_Programming #Two_Pointers #LeetCode_75_Two_Pointers
// #Dynamic_Programming_I_Day_19 #Level_1_Day_2_String #Udemy_Two_Pointers
// #Top_Interview_150_Two_Pointers #2025_05_24_Time_0_ms_(100.00%)_Space_3.90_MB_(76.98%)

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	n := len(t)
	m := len(s)
	if m == 0 {
		return true
	}
	for j < n {
		if s[i] == t[j] {
			i++
			if i == m {
				return true
			}
		}
		j++
	}
	return false
}
