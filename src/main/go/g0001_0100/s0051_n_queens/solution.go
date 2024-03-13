package s0051_n_queens

// #Hard #Top_100_Liked_Questions #Array #Backtracking #Big_O_Time_O(N!)_Space_O(N)
// #2024_03_13_Time_0_ms_(100.00%)_Space_3.3_MB_(61.83%)

func solveNQueens(n int) [][]string {
	pos := make([]bool, n+2*n-1+2*n-1)
	pos2 := make([]int, n)
	ans := [][]string{}
	helper(n, 0, pos, pos2, &ans)
	return ans
}

func helper(n, row int, pos []bool, pos2 []int, ans *[][]string) {
	if row == n {
		construct(n, pos2, ans)
		return
	}
	for i := 0; i < n; i++ {
		index := n + 2*n - 1 + n - 1 + i - row
		if pos[i] || pos[n+i+row] || pos[index] {
			continue
		}
		pos[i] = true
		pos[n+i+row] = true
		pos[index] = true
		pos2[row] = i
		helper(n, row+1, pos, pos2, ans)
		pos[i] = false
		pos[n+i+row] = false
		pos[index] = false
	}
}

func construct(n int, pos []int, ans *[][]string) {
	sol := []string{}
	for r := 0; r < n; r++ {
		queenRow := make([]byte, n)
		for i := range queenRow {
			queenRow[i] = '.'
		}
		queenRow[pos[r]] = 'Q'
		sol = append(sol, string(queenRow))
	}
	*ans = append(*ans, sol)
}
