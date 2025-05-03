package s0048_rotate_image

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Math #Matrix
// #Data_Structure_II_Day_3_Array #Programming_Skills_II_Day_7 #Udemy_2D_Arrays/Matrix
// #Top_Interview_150_Matrix #Big_O_Time_O(n^2)_Space_O(1)
// #2025_05_03_Time_0_ms_(100.00%)_Space_4.06_MB_(97.99%)

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			pos := [][]int{
				{i, j}, {j, n - 1 - i}, {n - 1 - i, n - 1 - j}, {n - 1 - j, i},
			}
			t := matrix[pos[0][0]][pos[0][1]]
			for k := 1; k < len(pos); k++ {
				temp := matrix[pos[k][0]][pos[k][1]]
				matrix[pos[k][0]][pos[k][1]] = t
				t = temp
			}
			matrix[pos[0][0]][pos[0][1]] = t
		}
	}
}
