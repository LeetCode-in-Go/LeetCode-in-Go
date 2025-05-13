package s0057_insert_interval

// #Medium #Array #Level_2_Day_17_Interval #Top_Interview_150_Intervals
// #2025_05_13_Time_0_ms_(100.00%)_Space_6.02_MB_(71.75%)

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	l := 0
	r := n - 1
	for l < n && newInterval[0] > intervals[l][1] {
		l++
	}
	for r >= 0 && newInterval[1] < intervals[r][0] {
		r--
	}
	res := make([][]int, l+n-r)
	for i := 0; i < l; i++ {
		res[i] = make([]int, 2)
		copy(res[i], intervals[i])
	}
	res[l] = make([]int, 2)
	if l == n {
		res[l][0] = newInterval[0]
	} else {
		res[l][0] = min(newInterval[0], intervals[l][0])
	}
	if r == -1 {
		res[l][1] = newInterval[1]
	} else {
		res[l][1] = max(newInterval[1], intervals[r][1])
	}
	for i, j := l+1, r+1; j < n; i, j = i+1, j+1 {
		res[i] = intervals[j]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
