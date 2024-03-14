package s0073_set_matrix_zeroes

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table #Matrix
// #Udemy_2D_Arrays/Matrix #Big_O_Time_O(m*n)_Space_O(1)
// #2024_03_14_Time_8_ms_(83.64%)_Space_5.9_MB_(75.58%)

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row0 := false
	col0 := false

	// Check if 0th col needs to be marked all 0s in future
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
			break
		}
	}

	// Check if 0th row needs to be marked all 0s in future
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	// Store the signals in 0th row and column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Mark 0 for all cells based on signal from 0th row and 0th column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Set 0th column
	for i := 0; i < m; i++ {
		if col0 {
			matrix[i][0] = 0
		}
	}

	// Set 0th row
	for i := 0; i < n; i++ {
		if row0 {
			matrix[0][i] = 0
		}
	}
}
