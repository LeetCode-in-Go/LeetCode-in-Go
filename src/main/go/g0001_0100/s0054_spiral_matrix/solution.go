package s0054_spiral_matrix

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Matrix #Simulation
// #Programming_Skills_II_Day_8 #Level_2_Day_1_Implementation/Simulation #Udemy_2D_Arrays/Matrix
// #Top_Interview_150_Matrix #2025_05_13_Time_0_ms_(100.00%)_Space_4.01_MB_(8.58%)

func spiralOrder(matrix [][]int) []int {
	list := make([]int, 0)
	r := 0
	c := 0
	bigR := len(matrix) - 1
	bigC := len(matrix[0]) - 1
	for r <= bigR && c <= bigC {
		for i := c; i <= bigC; i++ {
			list = append(list, matrix[r][i])
		}
		r++
		for i := r; i <= bigR; i++ {
			list = append(list, matrix[i][bigC])
		}
		bigC--
		for i := bigC; i >= c && r <= bigR; i-- {
			list = append(list, matrix[bigR][i])
		}
		bigR--
		for i := bigR; i >= r && c <= bigC; i-- {
			list = append(list, matrix[i][c])
		}
		c++
	}
	return list
}
