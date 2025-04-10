package s0049_group_anagrams

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #String #Hash_Table #Sorting
// #Data_Structure_II_Day_8_String #Programming_Skills_II_Day_11 #Udemy_Strings
// #Top_Interview_150_Hashmap #Big_O_Time_O(n*k_log_k)_Space_O(n)
// #2024_03_13_Time_16_ms_(85.07%)_Space_7.2_MB_(85.47%)

func groupAnagrams(strings []string) [][]string {
	stringsByPattern := make(map[[26]byte][]string)
	for i := 0; i < len(strings); i++ {
		var pattern [26]byte
		for c := 0; c < len(strings[i]); c++ {
			pattern[strings[i][c]-'a']++
		}
		stringsByPattern[pattern] = append(stringsByPattern[pattern], strings[i])
	}

	answer, index := make([][]string, len(stringsByPattern)), 0
	for _, slice := range stringsByPattern {
		answer[index] = slice
		index++
	}
	return answer
}
