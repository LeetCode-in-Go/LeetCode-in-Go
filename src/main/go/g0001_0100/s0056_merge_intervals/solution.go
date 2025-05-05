package s0056_merge_intervals

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Sorting
// #Data_Structure_II_Day_2_Array #Level_2_Day_17_Interval #Udemy_2D_Arrays/Matrix
// #Top_Interview_150_Intervals #Big_O_Time_O(n_log_n)_Space_O(n)
// #2025_05_05_Time_0_ms_(100.00%)_Space_8.61_MB_(25.59%)

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{intervals[0]}
	for i, j := 0, 1; j < len(intervals); j++ {
		if res[i][1] >= intervals[j][0] {
			res[i] = []int{res[i][0], max(res[i][1], intervals[j][1])}
		} else {
			res = append(res, intervals[j])
			i++
		}
	}

	return res
}
