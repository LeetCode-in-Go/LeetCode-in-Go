package s0399_evaluate_division

// #Medium #Array #Depth_First_Search #Breadth_First_Search #Graph #Union_Find #Shortest_Path
// #LeetCode_75_Graphs/DFS #Top_Interview_150_Graph_General
// #2025_05_24_Time_0_ms_(100.00%)_Space_4.07_MB_(91.71%)

type Solution struct {
	root map[string]string
	rate map[string]float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	sol := &Solution{
		root: make(map[string]string),
		rate: make(map[string]float64),
	}
	n := len(equations)
	for _, equation := range equations {
		x := equation[0]
		y := equation[1]
		sol.root[x] = x
		sol.root[y] = y
		sol.rate[x] = 1.0
		sol.rate[y] = 1.0
	}
	for i := 0; i < n; i++ {
		x := equations[i][0]
		y := equations[i][1]
		sol.union(x, y, values[i])
	}
	result := make([]float64, len(queries))
	for i, query := range queries {
		x := query[0]
		y := query[1]
		if _, exists := sol.root[x]; !exists {
			result[i] = -1.0
			continue
		}
		if _, exists := sol.root[y]; !exists {
			result[i] = -1.0
			continue
		}
		rootX := sol.findRoot(x, x, 1.0)
		rootY := sol.findRoot(y, y, 1.0)
		if rootX == rootY {
			result[i] = sol.rate[x] / sol.rate[y]
		} else {
			result[i] = -1.0
		}
	}
	return result
}

func (sol *Solution) union(x, y string, v float64) {
	rootX := sol.findRoot(x, x, 1.0)
	rootY := sol.findRoot(y, y, 1.0)
	sol.root[rootX] = rootY
	r1 := sol.rate[x]
	r2 := sol.rate[y]
	sol.rate[rootX] = v * r2 / r1
}

func (sol *Solution) findRoot(originalX, x string, r float64) string {
	if sol.root[x] == x {
		sol.root[originalX] = x
		sol.rate[originalX] = r * sol.rate[x]
		return x
	}
	return sol.findRoot(originalX, sol.root[x], r*sol.rate[x])
}
