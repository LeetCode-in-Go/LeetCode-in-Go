package s0013_roman_to_integer

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #String #Hash_Table #Math
// #Top_Interview_150_Array/String #2025_05_12_Time_0_ms_(100.00%)_Space_4.73_MB_(58.29%)

func romanToInt(s string) int {
	x := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			x = getX(s, x, i, 1, 'V', 'X')
		case 'V':
			x += 5
		case 'X':
			x = getX(s, x, i, 10, 'L', 'C')
		case 'L':
			x += 50
		case 'C':
			x = getX(s, x, i, 100, 'D', 'M')
		case 'D':
			x += 500
		case 'M':
			x += 1000
		}
	}
	return x
}

func getX(s string, x, i, i2 int, v, x2 byte) int {
	if i+1 == len(s) {
		x += i2
	} else if s[i+1] == v {
		x -= i2
	} else if s[i+1] == x2 {
		x -= i2
	} else {
		x += i2
	}
	return x
}
