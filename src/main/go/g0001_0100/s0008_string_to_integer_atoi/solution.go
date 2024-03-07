package s0008_string_to_integer_atoi

// #Medium #Top_Interview_Questions #String #2024_03_07_Time_0_ms_(100.00%)_Space_2.3_MB_(27.02%)

func myAtoi(s string) (num int) {
	amountToStrip := 0
	for _, c := range s {
		if c == ' ' {
			amountToStrip++
		} else {
			break
		}
	}
	s = s[amountToStrip:]
	neg := false
	if len(s) == 0 {
		return 0
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	inp := ""
	for _, c := range s {
		if int(c) < 48 || int(c) > 57 {
			break
		}
		inp += string(c)
	}
	amountToStrip = 0
	for _, c := range inp {
		if c == '0' {
			amountToStrip++
		} else {
			break
		}
	}
	inp = inp[amountToStrip:]
	for i, c := range inp {
		initial := num
		num += int(c) - 48
		if i != len(inp)-1 {
			num *= 10
		}
		if num < initial {
			num = 2147483649
			break
		}
	}

	if neg {
		num = 0 - num
	}
	if num > 2147483647 {
		return 2147483647
	} else if num < -2147483648 {
		return -2147483648
	}
	return
}
