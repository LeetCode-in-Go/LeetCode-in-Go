package s0032_longest_valid_parentheses

// #Hard #Top_100_Liked_Questions #String #Dynamic_Programming #Stack #Big_O_Time_O(n)_Space_O(1)
// #2025_04_27_Time_0_ms_(100.00%)_Space_4.28_MB_(96.40%)

func longestValidParentheses(s string) int {
	max := 0
	left, right := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch == '(' {
			left++
		} else {
			right++
		}
		if right > left {
			left, right = 0, 0
		}
		if left == right {
			max = maxInt(max, left+right)
		}
	}
	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		ch := s[i]
		if ch == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			left, right = 0, 0
		}
		if left == right {
			max = maxInt(max, left+right)
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
