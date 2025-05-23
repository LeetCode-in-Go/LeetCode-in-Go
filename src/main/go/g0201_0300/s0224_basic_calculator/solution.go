package s0224_basic_calculator

// #Hard #String #Math #Stack #Recursion #Top_Interview_150_Stack
// #2025_05_23_Time_0_ms_(100.00%)_Space_5.89_MB_(39.77%)

type Solution struct {
	i int
}

func calculate(s string) int {
	sol := &Solution{}
	return sol.helper([]rune(s))
}

func (sol *Solution) helper(ca []rune) int {
	num := 0
	prenum := 0
	isPlus := true
	for ; sol.i < len(ca); sol.i++ {
		c := ca[sol.i]
		if c != ' ' {
			if c >= '0' && c <= '9' {
				if num == 0 {
					num = int(c - '0')
				} else {
					num = num*10 + int(c-'0')
				}
			} else if c == '+' {
				prenum += num * sol.boolToInt(isPlus)
				isPlus = true
				num = 0
			} else if c == '-' {
				prenum += num * sol.boolToInt(isPlus)
				num = 0
				isPlus = false
			} else if c == '(' {
				sol.i++
				prenum += sol.helper(ca) * sol.boolToInt(isPlus)
				isPlus = true
				num = 0
			} else if c == ')' {
				return prenum + num*sol.boolToInt(isPlus)
			}
		}
	}
	return prenum + num*sol.boolToInt(isPlus)
}

func (sol *Solution) boolToInt(b bool) int {
	if b {
		return 1
	}
	return -1
}
