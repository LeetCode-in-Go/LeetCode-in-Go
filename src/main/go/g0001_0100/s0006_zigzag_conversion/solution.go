package s0006_zigzag_conversion

// #Medium #String #Top_Interview_150_Array/String
// #2024_03_05_Time_0_ms_(100.00%)_Space_4_MB_(86.83%)

func convert(s string, numRows int) string {
	topJump := (numRows-2)*2 + 2
	if topJump <= 0 {
		return s
	}
	size := len(s)
	resultArray := make([]byte, size)
	j := 0
	for i := 0; i < size; i += topJump {
		resultArray[j] = s[i]
		j++
	}
	for row := 1; row < numRows-1; row++ {
		for i := 0; i-row < size; i += topJump {
			iLeft := i - row
			iRight := i + row
			if iLeft > 0 {
				resultArray[j] = s[iLeft]
				j++
			}
			if iRight < size {
				resultArray[j] = s[iRight]
				j++
			}
		}
	}
	for i := topJump - numRows + 1; i < size; i += topJump {
		resultArray[j] = s[i]
		j++
	}
	return string(resultArray)
}
