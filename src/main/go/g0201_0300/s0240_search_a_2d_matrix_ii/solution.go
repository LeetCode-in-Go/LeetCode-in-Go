package s0240_search_a_2d_matrix_ii

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Binary_Search #Matrix
// #Divide_and_Conquer #Data_Structure_II_Day_4_Array #Binary_Search_II_Day_8
// #Big_O_Time_O(n+m)_Space_O(1) #2024_03_18_Time_11_ms_(96.10%)_Space_6.6_MB_(39.51%)

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
