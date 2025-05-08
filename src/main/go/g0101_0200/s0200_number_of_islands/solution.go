package s0200_number_of_islands

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Depth_First_Search
// #Breadth_First_Search #Matrix #Union_Find
// #Algorithm_II_Day_6_Breadth_First_Search_Depth_First_Search
// #Graph_Theory_I_Day_1_Matrix_Related_Problems #Level_1_Day_9_Graph/BFS/DFS #Udemy_Graph
// #Top_Interview_150_Graph_General #Big_O_Time_O(M*N)_Space_O(M*N)
// #2025_05_08_Time_2_ms_(87.44%)_Space_5.69_MB_(90.06%)

func numIslands(grid [][]byte) int {
	islands := 0
	if grid != nil && len(grid) != 0 && len(grid[0]) != 0 {
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == '1' {
					dfs(grid, i, j)
					islands++
				}
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || len(grid) <= x || y < 0 || len(grid[0]) <= y || grid[x][y] != '1' {
		return
	}
	grid[x][y] = 'x'
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
