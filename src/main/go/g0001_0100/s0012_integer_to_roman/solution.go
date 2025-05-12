package s0012_integer_to_roman

// #Medium #String #Hash_Table #Math #Top_Interview_150_Array/String
// #2025_05_12_Time_0_ms_(100.00%)_Space_4.94_MB_(97.25%)

import (
	"strings"
)

func intToRoman(num int) string {
	var result strings.Builder
	m := 1000
	c := 100
	x := 10
	i := 1

	num = numerals(&result, num, m, ' ', ' ', 'M')
	num = numerals(&result, num, c, 'M', 'D', 'C')
	num = numerals(&result, num, x, 'C', 'L', 'X')
	numerals(&result, num, i, 'X', 'V', 'I')

	return result.String()
}

func numerals(sb *strings.Builder, num, one int, cTen, cFive, cOne rune) int {
	div := num / one
	switch div {
	case 9:
		sb.WriteRune(cOne)
		sb.WriteRune(cTen)
	case 8:
		sb.WriteRune(cFive)
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
	case 7:
		sb.WriteRune(cFive)
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
	case 6:
		sb.WriteRune(cFive)
		sb.WriteRune(cOne)
	case 5:
		sb.WriteRune(cFive)
	case 4:
		sb.WriteRune(cOne)
		sb.WriteRune(cFive)
	case 3:
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
	case 2:
		sb.WriteRune(cOne)
		sb.WriteRune(cOne)
	case 1:
		sb.WriteRune(cOne)
	}
	return num - (div * one)
}
