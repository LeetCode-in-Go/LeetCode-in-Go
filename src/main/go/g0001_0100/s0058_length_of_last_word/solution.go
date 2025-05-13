package s0058_length_of_last_word

// #Easy #String #Programming_Skills_II_Day_6 #Udemy_Arrays #Top_Interview_150_Array/String
// #2025_05_13_Time_0_ms_(100.00%)_Space_4.00_MB_(100.00%)

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == ' ' && length > 0 {
			break
		} else if ch != ' ' {
			length++
		}
	}
	return length
}
