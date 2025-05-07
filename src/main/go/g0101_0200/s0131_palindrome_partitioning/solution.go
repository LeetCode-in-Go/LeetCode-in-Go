package s0131_palindrome_partitioning

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming
// #Backtracking #Big_O_Time_O(N*2^N)_Space_O(2^N*N)
// #2025_05_07_Time_18_ms_(92.42%)_Space_27.26_MB_(16.97%)

func partition(s string) [][]string {
	var res [][]string
	backtracking(&res, []string{}, s, 0)
	return res
}

func backtracking(res *[][]string, currArr []string, s string, start int) {
	if start == len(s) {
		*res = append(*res, append([]string(nil), currArr...))
		return
	}
	for end := start; end < len(s); end++ {
		if !isPalindrome(s, start, end) {
			continue
		}
		currArr = append(currArr, s[start:end+1])
		backtracking(res, currArr, s, end+1)
		currArr = currArr[:len(currArr)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
