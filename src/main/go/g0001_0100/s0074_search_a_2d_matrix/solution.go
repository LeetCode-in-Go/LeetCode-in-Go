package s0074_search_a_2d_matrix

// #Medium #Top_100_Liked_Questions #Array #Binary_Search #Matrix #Data_Structure_I_Day_5_Array
// #Algorithm_II_Day_1_Binary_Search #Binary_Search_I_Day_8 #Level_2_Day_8_Binary_Search
// #Udemy_2D_Arrays/Matrix #Top_Interview_150_Binary_Search #Big_O_Time_O(endRow+endCol)_Space_O(1)
// #2024_03_15_Time_0_ms_(100.00%)_Space_2.6_MB_(75.18%)

func searchMatrix(matrix [][]int, target int) bool {
	endRow := len(matrix)
	endCol := len(matrix[0])
	targetRow := 0
	result := false
	for i := 0; i < endRow; i++ {
		if matrix[i][endCol-1] >= target {
			targetRow = i
			break
		}
	}
	for i := 0; i < endCol; i++ {
		if matrix[targetRow][i] == target {
			result = true
			break
		}
	}
	return result
}
