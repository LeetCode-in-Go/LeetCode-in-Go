package s0452_minimum_number_of_arrows_to_burst_balloons

// #Medium #Array #Sorting #Greedy #LeetCode_75_Intervals #Top_Interview_150_Intervals
// #2025_05_24_Time_32_ms_(94.66%)_Space_16.72_MB_(69.38%)

import "slices"

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(acc []int, b []int) int {
		if acc[0] < b[0] {
			return -1
		} else if acc[0] > b[0] {
			return +1
		}
		return 0
	})
	openInterval := points[0]
	runningArrows := 1
	for i := 1; i < len(points); i++ {
		curr := points[i]
		// fmt.Printf("current: %v |  open interval %v\n", curr, openInterval)
		// intersects with the running interval?
		if curr[0] <= openInterval[1] {
			// we can use the same arrow!
			// let's adjust the running interval
			openInterval[0] = max(openInterval[0], curr[0])
			openInterval[1] = min(openInterval[1], curr[1])
		} else {
			runningArrows++
			openInterval = curr
		}
	}
	return runningArrows
}
