package s0071_simplify_path

// #Medium #String #Stack #Top_Interview_150_Stack
// #2025_05_14_Time_0_ms_(100.00%)_Space_5.75_MB_(22.88%)

func simplifyPath(path string) string {
	stk := make([]string, 0)
	start := 0
	for start < len(path) {
		for start < len(path) && path[start] == '/' {
			start++
		}
		end := start
		for end < len(path) && path[end] != '/' {
			end++
		}
		s := path[start:end]
		if s == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else if s != "." && s != "" {
			stk = append(stk, s)
		}
		start = end + 1
	}
	if len(stk) == 0 {
		return "/"
	}
	ans := ""
	for _, s := range stk {
		ans += "/" + s
	}
	return ans
}
