package s0210_course_schedule_ii

// #Medium #Top_Interview_Questions #Depth_First_Search #Breadth_First_Search #Graph
// #Topological_Sort #Level_2_Day_11_Graph/BFS/DFS #Top_Interview_150_Graph_General
// #2025_05_23_Time_0_ms_(100.00%)_Space_7.44_MB_(79.02%)

func findOrder(numCourses int, edges [][]int) []int {
	courses := make([][]int, numCourses)
	for _, edge := range edges {
		if courses[edge[0]] == nil {
			courses[edge[0]] = make([]int, 0)
		}
		courses[edge[0]] = append(courses[edge[0]], edge[1])
	}
	visitMatrix := make([]byte, numCourses)
	ans := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if !addOrder(courses, visitMatrix, i, &ans) {
			return []int{}
		}
	}
	return ans
}

func addOrder(courses [][]int, visits []byte, start int, ans *[]int) bool {
	if visits[start] == 1 {
		return false
	}
	if visits[start] == 2 {
		return true
	}
	visits[start] = 1
	for _, next := range courses[start] {
		if !addOrder(courses, visits, next, ans) {
			return false
		}
	}
	*ans = append(*ans, start)
	visits[start] = 2
	return true
}
