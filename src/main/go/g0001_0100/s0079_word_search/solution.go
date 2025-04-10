package s0079_word_search

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Matrix #Backtracking
// #Algorithm_II_Day_11_Recursion_Backtracking #Top_Interview_150_Backtracking
// #Big_O_Time_O(4^(m*n))_Space_O(m*n) #2024_03_15_Time_103_ms_(79.70%)_Space_2_MB_(58.16%)

func exist(board [][]byte, word string) bool {
	var run func(i, j, k int) bool
	run = func(i, j, k int) bool {
		if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
			return false
		}

		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		board[i][j] = ' '
		results := run(i+1, j, k+1) ||
			run(i-1, j, k+1) ||
			run(i, j+1, k+1) ||
			run(i, j-1, k+1)

		board[i][j] = word[k]

		return results
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if run(i, j, 0) {
				return true
			}
		}
	}

	return false
}
