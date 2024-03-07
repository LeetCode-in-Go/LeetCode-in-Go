package s0010_regular_expression_matching

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #String #Dynamic_Programming #Recursion
// #Udemy_Dynamic_Programming #Big_O_Time_O(m*n)_Space_O(m*n)
// #2024_03_07_Time_0_ms_(100.00%)_Space_2.3_MB_(36.02%)

func isMatch(s string, p string) bool {
	textLength := len(s)
	patternLength := len(p)
	result := make([][]bool, textLength+1)
	for i := range result {
		result[i] = make([]bool, patternLength+1)
	}
	result[0][0] = true
	for j := 2; j < patternLength+1; {
		if p[j-1] == '*' {
			result[0][j] = result[0][j-2]
		}
		j++
	}
	for i := 0; i < textLength; {
		for j := 0; j < patternLength; {
			if p[j] == s[i] || p[j] == '.' {
				result[i+1][j+1] = result[i][j]
			} else if p[j] == '*' {
				if p[j-1] == s[i] || p[j-1] == '.' {
					result[i+1][j+1] = result[i+1][j-1] || result[i][j+1]
				} else {
					result[i+1][j+1] = result[i+1][j-1]
				}
			} else {
				result[i+1][j+1] = false
			}
			j++
		}
		i++
	}
	return result[textLength][patternLength]
}
