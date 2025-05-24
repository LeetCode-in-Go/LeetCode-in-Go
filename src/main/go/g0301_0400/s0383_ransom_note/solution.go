package s0383_ransom_note

// #Easy #String #Hash_Table #Counting #Data_Structure_I_Day_6_String #Top_Interview_150_Hashmap
// #2025_05_24_Time_0_ms_(100.00%)_Space_5.76_MB_(86.24%)

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make([]int, 26)
	n := len(ransomNote)
	for i := 0; i < n; i++ {
		charCount[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine) && n != 0; i++ {
		if charCount[magazine[i]-'a'] > 0 {
			n--
			charCount[magazine[i]-'a']--
		}
	}
	return n == 0
}
