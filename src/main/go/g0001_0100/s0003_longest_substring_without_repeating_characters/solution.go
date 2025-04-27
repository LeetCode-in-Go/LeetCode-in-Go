package s0003_longest_substring_without_repeating_characters

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Sliding_Window
// #Algorithm_I_Day_6_Sliding_Window #Level_2_Day_14_Sliding_Window/Two_Pointer #Udemy_Strings
// #Top_Interview_150_Sliding_Window #Big_O_Time_O(n)_Space_O(1) #AI_can_be_used_to_solve_the_task
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.45_MB_(94.51%)

func lengthOfLongestSubstring(s string) int {
	lastIndices := make([]int, 256)
	for i := 0; i < 256; i++ {
		lastIndices[i] = -1
	}
	maxLen := 0
	curLen := 0
	start := 0

	for i := 0; i < len(s); i++ {
		cur := s[i]
		if lastIndices[cur] < start {
			lastIndices[cur] = i
			curLen++
		} else {
			lastIndex := lastIndices[cur]
			start = lastIndex + 1
			curLen = i - start + 1
			lastIndices[cur] = i
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
