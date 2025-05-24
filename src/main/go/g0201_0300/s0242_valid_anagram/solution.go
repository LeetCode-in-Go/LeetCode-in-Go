package s0242_valid_anagram

// #Easy #String #Hash_Table #Sorting #Data_Structure_I_Day_6_String
// #Programming_Skills_I_Day_11_Containers_and_Libraries #Udemy_Strings #Top_Interview_150_Hashmap
// #2025_05_24_Time_0_ms_(100.00%)_Space_4.81_MB_(36.79%)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charFreqMap := make([]int, 26)
	for _, c := range s {
		charFreqMap[c-'a']++
	}
	for _, c := range t {
		if charFreqMap[c-'a'] == 0 {
			return false
		}
		charFreqMap[c-'a']--
	}
	return true
}
