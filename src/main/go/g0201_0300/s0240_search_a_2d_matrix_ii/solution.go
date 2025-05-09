package s0240_search_a_2d_matrix_ii

// #Medium #Top_100_Liked_Questions #Array #Binary_Search #Matrix #Divide_and_Conquer
// #Data_Structure_II_Day_4_Array #Binary_Search_II_Day_8 #Big_O_Time_O(n+m)_Space_O(1)
// #2025_05_10_Time_9_ms_(98.86%)_Space_8.76_MB_(12.12%)

func searchMatrix(matrix [][]int, target int) bool {
	r := 0
	c := len(matrix[0]) - 1
	for r < len(matrix) && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}
