package s0207_course_schedule

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search
// #Breadth_First_Search #Graph #Topological_Sort #Top_Interview_150_Graph_General
// #Big_O_Time_O(N)_Space_O(N) #2025_05_09_Time_0_ms_(100.00%)_Space_7.98_MB_(72.75%)

type State int

const (
	Unvisited State = iota
	Visiting
	Visited
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]State, numCourses)
	graph := make([][]int, numCourses)
	for _, dep := range prerequisites {
		graph[dep[1]] = append(graph[dep[1]], dep[0])
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == Unvisited {
			if !dfs(i, visited, graph) {
				return false
			}
		}
	}
	return true
}

func dfs(start int, visited []State, graph [][]int) bool {
	if visited[start] == Visiting {
		return false
	}
	if visited[start] == Visited {
		return true
	}
	visited[start] = Visiting
	for _, next := range graph[start] {
		if !dfs(next, visited, graph) {
			return false
		}
	}
	visited[start] = Visited
	return true
}
