package s0017_letter_combinations_of_a_phone_number

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Backtracking
// #LeetCode_75_Backtracking #Algorithm_II_Day_11_Recursion_Backtracking
// #Udemy_Backtracking/Recursion #Top_Interview_150_Backtracking #Big_O_Time_O(4^n)_Space_O(n)
// #2025_04_27_Time_0_ms_(100.00%)_Space_3.96_MB_(93.28%)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var ans []string
	var sb string
	findCombinations(0, digits, letters, &sb, &ans)
	return ans
}

func findCombinations(start int, nums string, letters []string, curr *string, ans *[]string) {
	if len(*curr) == len(nums) {
		*ans = append(*ans, *curr)
		return
	}
	for i := start; i < len(nums); i++ {
		n := int(nums[i] - '0')
		for j := 0; j < len(letters[n]); j++ {
			ch := letters[n][j]
			*curr += string(ch)
			findCombinations(i+1, nums, letters, curr, ans)
			*curr = (*curr)[:len(*curr)-1]
		}
	}
}
