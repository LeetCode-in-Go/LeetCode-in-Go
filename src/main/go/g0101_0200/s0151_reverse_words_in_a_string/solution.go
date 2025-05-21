package s0151_reverse_words_in_a_string

// #Medium #String #Two_Pointers #LeetCode_75_Array/String #Udemy_Strings
// #Top_Interview_150_Array/String #2025_05_21_Time_0_ms_(100.00%)_Space_5.37_MB_(51.24%)

func reverseWords(s string) string {
	var result []rune
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			i--
			continue
		}
		start := i
		for start >= 0 && s[start] != ' ' {
			start--
		}
		if len(result) > 0 {
			result = append(result, ' ')
		}
		result = append(result, []rune(s[start+1:i+1])...)
		i = start - 1
	}
	return string(result)
}
