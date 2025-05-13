package s0036_valid_sudoku

// #Medium #Top_Interview_Questions #Array #Hash_Table #Matrix #Data_Structure_I_Day_5_Array
// #Top_Interview_150_Matrix #2025_05_13_Time_0_ms_(100.00%)_Space_4.66_MB_(82.08%)

type Solution struct {
	j1 int
	i1 [9]int
	b1 [9]int
}

func isValidSudoku(board [][]byte) bool {
	s := &Solution{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.checkValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func (s *Solution) checkValid(board [][]byte, i, j int) bool {
	if j == 0 {
		s.j1 = 0
	}
	if board[i][j] == '.' {
		return true
	}
	val := int(board[i][j] - '0')
	if s.j1 == (s.j1 | (1 << (val - 1))) {
		return false
	}
	s.j1 |= 1 << (val - 1)
	if s.i1[j] == (s.i1[j] | (1 << (val - 1))) {
		return false
	}
	s.i1[j] |= 1 << (val - 1)
	b := (i/3)*3 + j/3
	if s.b1[b] == (s.b1[b] | (1 << (val - 1))) {
		return false
	}
	s.b1[b] |= 1 << (val - 1)
	return true
}
