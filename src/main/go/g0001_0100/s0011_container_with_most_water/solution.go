package s0011_container_with_most_water

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Greedy #Two_Pointers
// #LeetCode_75_Two_Pointers #Algorithm_II_Day_4_Two_Pointers #Top_Interview_150_Two_Pointers
// #Big_O_Time_O(n)_Space_O(1) #2025_04_27_Time_0_ms_(100.00%)_Space_9.30_MB_(86.74%)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > maxArea {
			maxArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
