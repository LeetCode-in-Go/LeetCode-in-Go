package s0739_daily_temperatures

// #Medium #Top_100_Liked_Questions #Array #Stack #Monotonic_Stack #Programming_Skills_II_Day_6
// #Big_O_Time_O(n)_Space_O(n) #2024_03_10_Time_120_ms_(98.07%)_Space_8.2_MB_(98.62%)

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
