package s0009_palindrome_number

// #Easy #Math #Udemy_Integers #Top_Interview_150_Math
// #2025_04_27_Time_0_ms_(100.00%)_Space_5.96_MB_(99.06%)

func isPalindrome(x int) bool {
	if x < 10 {
		return x >= 0
	}
	if x%10 == 0 {
		return false
	}
	a, b := x, 0
	for a > b {
		a, b = a/10, b*10+a%10
	}
	return a == b || a == b/10
}
