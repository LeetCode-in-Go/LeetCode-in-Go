package s0289_game_of_life

// #Medium #Array #Matrix #Simulation #Top_Interview_150_Matrix
// #2025_05_24_Time_0_ms_(100.00%)_Space_3.96_MB_(66.98%)

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives := countLives(board, i, j, m, n)
			if board[i][j] == 0 && lives == 3 {
				board[i][j] = 2
			} else if board[i][j] == 1 && (lives == 2 || lives == 3) {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func countLives(board [][]int, i, j, m, n int) int {
	lives := 0
	for r := max(0, i-1); r <= min(m-1, i+1); r++ {
		for c := max(0, j-1); c <= min(n-1, j+1); c++ {
			lives += board[r][c] & 1
		}
	}
	lives -= board[i][j] & 1
	return lives
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
