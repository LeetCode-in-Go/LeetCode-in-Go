package s0008_string_to_integer_atoi

// #Medium #Top_Interview_Questions #String #2024_03_07_Time_0_ms_(100.00%)_Space_2.3_MB_(27.02%)

import (
	"math"
	"strings"
)

func myAtoi(A string) int {
	isPositive := checkSign(A)
	validStr := clear(A)
	n := len(validStr)
	if n == 0 {
		return 0
	}
	var result int64
	for i := 0; i < n; i++ {
		result = result*int64(10) + int64(validStr[i]-48)
		if isPositive && result >= math.MaxInt32 {
			return math.MaxInt32
		} else if !isPositive && result >= math.MaxInt32+1 {
			return math.MinInt32
		}
	}
	if !isPositive {
		return int(result) * -1
	}
	return int(result)
}

func checkSign(str string) bool {
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == '-' {
			return false
		} else {
			break
		}
	}
	return true
}

func clear(str string) string {
	var result strings.Builder
	n := len(str)
	signCount := 0
	for i := 0; i < n; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			result.WriteByte(str[i])
		} else if (str[i] == '-' || str[i] == '+') && signCount < 1 && result.Len() == 0 {
			signCount++
			continue
		} else if str[i] == ' ' && result.Len() == 0 && signCount == 0 {
			continue
		} else {
			break
		}
	}
	return result.String()
}
