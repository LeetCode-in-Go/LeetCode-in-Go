package s0022_generate_parentheses

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming
// #Backtracking #Algorithm_II_Day_11_Recursion_Backtracking #Udemy_Backtracking/Recursion
// #Top_Interview_150_Backtracking #Big_O_Time_O(2^n)_Space_O(n)
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.66_MB_(78.73%)

func generateParenthesis(n int) []string {
	result := []string{}
	generate("", 0, 0, n, &result)
	return result
}

func generate(s string, start, close, n int, result *[]string) {
	if len(s) == 2*n {
		*result = append(*result, s)
		return
	}
	if start < n {
		generate(s+"(", start+1, close, n, result)
	}
	if close < start {
		generate(s+")", start, close+1, n, result)
	}
}
