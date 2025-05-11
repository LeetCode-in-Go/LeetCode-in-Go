package s0438_find_all_anagrams_in_a_string

// #Medium #Top_100_Liked_Questions #String #Hash_Table #Sliding_Window
// #Algorithm_II_Day_5_Sliding_Window #Programming_Skills_II_Day_12
// #Level_1_Day_12_Sliding_Window/Two_Pointer #Big_O_Time_O(n+m)_Space_O(1)
// #2025_05_11_Time_0_ms_(100.00%)_Space_6.83_MB_(90.16%)

func findAnagrams(s string, p string) []int {
	var mapArr [26]int
	for i := 0; i < len(p); i++ {
		mapArr[p[i]-'a']++
	}
	var res []int
	i, j := 0, 0
	for i < len(s) {
		idx := s[i] - 'a'
		mapArr[idx]--
		if i >= len(p) {
			mapArr[s[j]-'a']++
			j++
		}
		finish := true
		for k := 0; k < 26; k++ {
			if mapArr[k] != 0 {
				finish = false
				break
			}
		}
		if i >= len(p)-1 && finish {
			res = append(res, j)
		}
		i++
	}
	return res
}
