package s0052_n_queens_ii

// #Hard #Backtracking #Top_Interview_150_Backtracking
// #2025_05_24_Time_0_ms_(100.00%)_Space_3.86_MB_(84.00%)

func totalNQueens(n int) int {
	row := make([]bool, n)
	col := make([]bool, n)
	diagonal := make([]bool, 2*n-1)
	antiDiagonal := make([]bool, 2*n-1)
	return totalNQueensHelper(n, 0, row, col, diagonal, antiDiagonal)
}

func totalNQueensHelper(
	n, r int,
	row, col, diagonal, antiDiagonal []bool,
) int {
	if r == n {
		return 1
	}
	count := 0
	for c := 0; c < n; c++ {
		if !row[r] && !col[c] && !diagonal[r+c] && !antiDiagonal[r-c+n-1] {
			row[r], col[c], diagonal[r+c], antiDiagonal[r-c+n-1] = true, true, true, true
			count += totalNQueensHelper(n, r+1, row, col, diagonal, antiDiagonal)
			row[r], col[c], diagonal[r+c], antiDiagonal[r-c+n-1] = false, false, false, false
		}
	}
	return count
}
