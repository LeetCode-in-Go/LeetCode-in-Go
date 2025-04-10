package s0064_minimum_path_sum

// #Medium #Top_100_Liked_Questions #Array #Dynamic_Programming #Matrix
// #Dynamic_Programming_I_Day_16 #Udemy_Dynamic_Programming #Top_Interview_150_Multidimensional_DP
// #Big_O_Time_O(m*n)_Space_O(m*n) #2024_03_14_Time_0_ms_(100.00%)_Space_4.1_MB_(24.44%)

func minPathSum(grid [][]int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	dm := make([][]int, len(grid))
	for i := range dm {
		dm[i] = make([]int, len(grid[0]))
	}
	sum := 0
	for r := len(grid) - 1; r >= 0; r-- {
		dm[r][len(grid[0])-1] = grid[r][len(grid[0])-1] + sum
		sum += grid[r][len(grid[0])-1]
	}
	sum = 0
	for c := len(grid[0]) - 1; c >= 0; c-- {
		dm[len(grid)-1][c] = grid[len(grid)-1][c] + sum
		sum += grid[len(grid)-1][c]
	}
	return recur(grid, dm, 0, 0)
}

func recur(grid [][]int, dm [][]int, r, c int) int {
	if dm[r][c] == 0 && r != len(grid)-1 && c != len(grid[0])-1 {
		dm[r][c] = grid[r][c] + min(recur(grid, dm, r+1, c), recur(grid, dm, r, c+1))
	}
	return dm[r][c]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
