package s0125_valid_palindrome

// #Easy #Top_Interview_Questions #String #Two_Pointers #Udemy_Two_Pointers
// #Top_Interview_150_Two_Pointers #2025_05_20_Time_0_ms_(100.00%)_Space_4.50_MB_(100.00%)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	res := true
	for res {
		for i < j && isNotAlphaNumeric(s[i]) {
			i++
		}
		for i < j && isNotAlphaNumeric(s[j]) {
			j--
		}
		if i >= j {
			break
		}
		left := upperToLower(s[i])
		right := upperToLower(s[j])
		if left != right {
			res = false
		}
		i++
		j--
	}
	return res
}

func isNotAlphaNumeric(c byte) bool {
	return (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9')
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func upperToLower(c byte) byte {
	if isUpper(c) {
		c += 32
	}
	return c
}
