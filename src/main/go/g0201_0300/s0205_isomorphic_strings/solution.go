package s0205_isomorphic_strings

// #Easy #String #Hash_Table #Level_1_Day_2_String #Top_Interview_150_Hashmap
// #2025_05_23_Time_0_ms_(100.00%)_Space_5.10_MB_(17.24%)

func isIsomorphic(s string, t string) bool {
	charMap := make([]int, 128)
	str := []rune(s)
	tar := []rune(t)
	n := len(str)
	for i := 0; i < n; i++ {
		if charMap[tar[i]] == 0 {
			if search(charMap, str[i], tar[i]) != -1 {
				return false
			}
			charMap[tar[i]] = int(str[i])
		} else {
			if charMap[tar[i]] != int(str[i]) {
				return false
			}
		}
	}
	return true
}

func search(charMap []int, tar rune, skip rune) int {
	for i := 0; i < 128; i++ {
		if i == int(skip) {
			continue
		}
		if charMap[i] != 0 && charMap[i] == int(tar) {
			return i
		}
	}
	return -1
}
