package s0063_unique_paths_ii

// #Medium #Array #Dynamic_Programming #Matrix #Dynamic_Programming_I_Day_15
// #Top_Interview_150_Multidimensional_DP #2025_05_14_Time_0_ms_(100.00%)_Space_4.14_MB_(96.06%)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = 1
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
		} else {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
