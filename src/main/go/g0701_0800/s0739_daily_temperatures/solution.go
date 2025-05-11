package s0739_daily_temperatures

// #Medium #Top_100_Liked_Questions #Array #Stack #Monotonic_Stack #LeetCode_75_Monotonic_Stack
// #Programming_Skills_II_Day_6 #Big_O_Time_O(n)_Space_O(n)
// #2025_05_11_Time_5_ms_(96.26%)_Space_10.49_MB_(86.49%)

func dailyTemperatures(temperatures []int) []int {
	sol := make([]int, len(temperatures))
	sol[len(temperatures)-1] = 0

	for i := len(sol) - 2; i >= 0; i-- {
		j := i + 1
		for j < len(sol) {
			if temperatures[i] < temperatures[j] {
				sol[i] = j - i
				break
			} else {
				if sol[j] == 0 {
					break
				}
				j += sol[j]
			}
		}
	}
	return sol
}
