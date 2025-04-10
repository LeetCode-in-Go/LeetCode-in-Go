package s0042_trapping_rain_water

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming #Two_Pointers
// #Stack #Monotonic_Stack #Dynamic_Programming_I_Day_9 #Udemy_Two_Pointers
// #Top_Interview_150_Array/String #Big_O_Time_O(n)_Space_O(1)
// #2024_03_12_Time_3_ms_(99.42%)_Space_5.4_MB_(97.01%)

func trap(height []int) int {
	if height == nil {
		return 0
	}
	l, r := 0, len(height)-1
	maxLeft, maxRight := height[l], height[r]
	res := 0
	for l < r {
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, height[l])
			res += maxLeft - height[l]
		} else {
			r--
			maxRight = max(maxRight, height[r])
			res += maxRight - height[r]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
