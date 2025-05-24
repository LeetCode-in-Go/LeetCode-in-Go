package s0433_minimum_genetic_mutation

// #Medium #String #Hash_Table #Breadth_First_Search #Graph_Theory_I_Day_12_Breadth_First_Search
// #Top_Interview_150_Graph_BFS #2025_05_24_Time_0_ms_(100.00%)_Space_3.94_MB_(50.69%)

func isInBank(set map[string]bool, cur string) []string {
	var res []string
	for each := range set {
		diff := 0
		for i := 0; i < len(each); i++ {
			if each[i] != cur[i] {
				diff++
				if diff > 1 {
					break
				}
			}
		}
		if diff == 1 {
			res = append(res, each)
		}
	}
	return res
}

func minMutation(start string, end string, bank []string) int {
	set := make(map[string]bool)
	for _, s := range bank {
		set[s] = true
	}
	queue := []string{start}
	step := 0
	for len(queue) > 0 {
		curSize := len(queue)
		for curSize > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur == end {
				return step
			}
			for _, next := range isInBank(set, cur) {
				queue = append(queue, next)
				delete(set, next)
			}
			curSize--
		}
		step++
	}
	return -1
}
