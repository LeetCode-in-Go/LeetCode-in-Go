package s0070_climbing_stairs

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Dynamic_Programming #Math #Memoization
// #Algorithm_I_Day_12_Dynamic_Programming #Dynamic_Programming_I_Day_2
// #Level_1_Day_10_Dynamic_Programming #Udemy_Dynamic_Programming #Top_Interview_150_1D_DP
// #Big_O_Time_O(n)_Space_O(n) #2025_05_05_Time_0_ms_(100.00%)_Space_3.88_MB_(81.88%)

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 2
	if n > 2 {
		for i := 2; i < n; i++ {
			f[i] = f[i-1] + f[i-2]
		}
	}
	return f[n-1]
}
