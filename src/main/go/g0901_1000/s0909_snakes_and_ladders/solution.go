package s0909_snakes_and_ladders

// #Medium #Array #Breadth_First_Search #Matrix #Top_Interview_150_Graph_BFS
// #2025_05_24_Time_0_ms_(100.00%)_Space_6.41_MB_(90.12%)

func snakesAndLadders(board [][]int) int {
	size := len(board)
	target := size * size
	visited := make([]bool, target)
	queue := []int{1}
	visited[0] = true
	step := 0

	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			previousLabel := queue[0]
			queue = queue[1:]

			if previousLabel == target {
				return step
			}

			for currentLabel := previousLabel + 1; currentLabel <= min(target, previousLabel+6); currentLabel++ {
				if visited[currentLabel-1] {
					continue
				}
				visited[currentLabel-1] = true
				position := indexToPosition(currentLabel, size)
				if board[position[0]][position[1]] == -1 {
					queue = append(queue, currentLabel)
				} else {
					queue = append(queue, board[position[0]][position[1]])
				}
			}
		}
		step++
	}

	return -1
}

func indexToPosition(index, size int) [2]int {
	vertical := size - 1 - (index-1)/size
	var horizontal int
	if (size-vertical)%2 == 1 {
		horizontal = (index - 1) % size
	} else {
		horizontal = size - 1 - (index-1)%size
	}
	return [2]int{vertical, horizontal}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
