package s0290_word_pattern

// #Easy #String #Hash_Table #Data_Structure_II_Day_7_String #Top_Interview_150_Hashmap
// #2025_05_24_Time_0_ms_(100.00%)_Space_4.07_MB_(21.29%)

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)
	for i, char := range pattern {
		word := words[i]
		if mappedWord, exists := charToWord[char]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			if _, exists := wordToChar[word]; exists {
				return false
			}
			charToWord[char] = word
			wordToChar[word] = char
		}
	}
	return true
}
