package s0084_largest_rectangle_in_histogram

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Array #Stack #Monotonic_Stack
// #Big_O_Time_O(n_log_n)_Space_O(log_n) #2024_03_15_Time_95_ms_(77.81%)_Space_7.8_MB_(91.91%)

import "math"

func largestRectangleArea(heights []int) int {
	return largestArea(heights, 0, len(heights))
}

func largestArea(a []int, start, limit int) int {
	if len(a) == 0 {
		return 0
	}
	if start == limit {
		return 0
	}
	if limit-start == 1 {
		return a[start]
	}
	if limit-start == 2 {
		maxOfTwoBars := max(a[start], a[start+1])
		areaFromTwo := min(a[start], a[start+1]) * 2
		return max(maxOfTwoBars, areaFromTwo)
	}
	if checkIfSorted(a, start, limit) {
		maxWhenSorted := 0
		for i := start; i < limit; i++ {
			if a[i]*(limit-i) > maxWhenSorted {
				maxWhenSorted = a[i] * (limit - i)
			}
		}
		return maxWhenSorted
	} else {
		minInd := findMinInArray(a, start, limit)
		return maxOfThreeNums(
			largestArea(a, start, minInd),
			a[minInd]*(limit-start),
			largestArea(a, minInd+1, limit),
		)
	}
}

func findMinInArray(a []int, start, limit int) int {
	min := math.MaxInt32
	minIndex := -1
	for index := start; index < limit; index++ {
		if a[index] < min {
			min = a[index]
			minIndex = index
		}
	}
	return minIndex
}

func checkIfSorted(a []int, start, limit int) bool {
	for i := start + 1; i < limit; i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxOfThreeNums(x, y, z int) int {
	return max(max(x, y), z)
}
