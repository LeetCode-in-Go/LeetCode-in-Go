package s0006_zigzag_conversion

// #Medium #String #2024_03_05_Time_0_ms_(100.00%)_Space_3.9_MB_(94.33%)

import (
	"strings"
)

func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows == 1 {
		return s
	}
	maxDist := numRows*2 - 2
	var buf strings.Builder
	for i := 0; i < numRows; i++ {
		index := i
		if i == 0 || i == numRows-1 {
			for index < sLen {
				buf.WriteByte(s[index])
				index += maxDist
			}
		} else {
			for index < sLen {
				buf.WriteByte(s[index])
				index += maxDist - i*2
				if index >= sLen {
					break
				}
				buf.WriteByte(s[index])
				index += i * 2
			}
		}
	}
	return buf.String()
}
