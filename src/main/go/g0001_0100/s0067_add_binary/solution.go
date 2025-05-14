package s0067_add_binary

// #Easy #String #Math #Bit_Manipulation #Simulation #Programming_Skills_II_Day_5
// #Top_Interview_150_Bit_Manipulation #2025_05_14_Time_0_ms_(100.00%)_Space_4.23_MB_(79.53%)

import (
	"strings"
)

func addBinary(a string, b string) string {
	var sb strings.Builder
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		sb.WriteByte(byte(sum%2 + '0'))
		carry = sum / 2
	}
	if carry != 0 {
		sb.WriteByte(byte(carry + '0'))
	}
	result := []byte(sb.String())
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
