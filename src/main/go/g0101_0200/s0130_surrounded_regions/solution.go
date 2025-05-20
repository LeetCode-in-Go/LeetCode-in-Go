package s0130_surrounded_regions

// #Medium #Top_Interview_Questions #Array #Depth_First_Search #Breadth_First_Search #Matrix
// #Union_Find #Algorithm_II_Day_8_Breadth_First_Search_Depth_First_Search
// #Top_Interview_150_Graph_General #2025_05_20_Time_0_ms_(100.00%)_Space_8.08_MB_(93.39%)

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			dfs(board, len(board)-1, i)
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, column int) {
	if row < 0 || row >= len(board) || column < 0 || column >= len(board[0]) || board[row][column] != 'O' {
		return
	}
	board[row][column] = '#'
	dfs(board, row+1, column)
	dfs(board, row-1, column)
	dfs(board, row, column+1)
	dfs(board, row, column-1)
}
