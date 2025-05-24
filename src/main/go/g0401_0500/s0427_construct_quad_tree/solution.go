package s0427_construct_quad_tree

// #Medium #Array #Tree #Matrix #Divide_and_Conquer #Top_Interview_150_Divide_and_Conquer
// #2025_05_24_Time_0_ms_(100.00%)_Space_8.07_MB_(80.28%)

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	return optimizedDfs(grid, 0, 0, len(grid))
}

func optimizedDfs(grid [][]int, rowStart, colStart, len int) *Node {
	zeroCount := 0
	oneCount := 0
	for row := rowStart; row < rowStart+len; row++ {
		isBreak := false
		for col := colStart; col < colStart+len; col++ {
			if grid[row][col] == 0 {
				zeroCount++
			} else {
				oneCount++
			}
			if oneCount > 0 && zeroCount > 0 {
				// We really no need to scan all.
				// Once we know there are both 0 and 1 then we can break.
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
	}
	if oneCount > 0 && zeroCount > 0 {
		midLen := len / 2
		topLeft := optimizedDfs(grid, rowStart, colStart, midLen)
		topRight := optimizedDfs(grid, rowStart, colStart+midLen, midLen)
		bottomLeft := optimizedDfs(grid, rowStart+midLen, colStart, midLen)
		bottomRight := optimizedDfs(grid, rowStart+midLen, colStart+midLen, midLen)
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     topLeft,
			TopRight:    topRight,
			BottomLeft:  bottomLeft,
			BottomRight: bottomRight,
		}
	} else {
		resultVal := oneCount > 0
		return &Node{
			Val:    resultVal,
			IsLeaf: true,
		}
	}
}
