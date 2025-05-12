package s0014_longest_common_prefix

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #String #Level_2_Day_2_String
// #Udemy_Strings #Top_Interview_150_Array/String
// #2025_05_12_Time_0_ms_(100.00%)_Space_4.28_MB_(73.97%)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	temp := strs[0]
	i := 1
	for temp != "" && i < len(strs) {
		if len(temp) > len(strs[i]) {
			temp = temp[:len(strs[i])]
		}
		cur := strs[i][:len(temp)]
		if cur != temp {
			temp = temp[:len(temp)-1]
		} else {
			i++
		}
	}
	return temp
}
