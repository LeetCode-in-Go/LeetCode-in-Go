package s0150_evaluate_reverse_polish_notation

// #Medium #Top_Interview_Questions #Array #Math #Stack #Programming_Skills_II_Day_3
// #Top_Interview_150_Stack #2025_05_21_Time_0_ms_(100.00%)_Space_6.10_MB_(53.08%)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if len(token) == 1 && !isDigit(token[0]) {
			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, eval(first, second, token))
		} else {
			num := 0
			sign := 1
			i := 0
			if token[0] == '-' {
				sign = -1
				i = 1
			}
			for ; i < len(token); i++ {
				num = num*10 + int(token[i]-'0')
			}
			stack = append(stack, num*sign)
		}
	}
	return stack[0]
}

func eval(first, second int, operator string) int {
	switch operator {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	default:
		return first / second
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
