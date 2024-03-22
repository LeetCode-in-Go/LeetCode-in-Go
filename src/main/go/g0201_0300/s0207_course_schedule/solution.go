package s0207_course_schedule

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search
// #Breadth_First_Search #Graph #Topological_Sort #Big_O_Time_O(N)_Space_O(N)
// #2024_03_22_Time_7_ms_(83.17%)_Space_6.3_MB_(51.07%)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	indegree := make([]int, numCourses)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		for _, v := range graph[i] {
			indegree[v]++
		}
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	visited := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited++
		for _, v := range graph[node] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return visited == numCourses
}
