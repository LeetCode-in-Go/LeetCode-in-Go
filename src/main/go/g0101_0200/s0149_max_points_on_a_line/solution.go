package s0149_max_points_on_a_line

// #Hard #Top_Interview_Questions #Array #Hash_Table #Math #Geometry #Algorithm_II_Day_21_Others
// #Top_Interview_150_Math #2025_05_24_Time_4_ms_(100.00%)_Space_4.11_MB_(93.17%)

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	max := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			c := x*points[j][1] - y*points[j][0]
			count := 2
			for k := j + 1; k < len(points); k++ {
				if c == x*points[k][1]-y*points[k][0] {
					count++
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
