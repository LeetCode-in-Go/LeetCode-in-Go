package s0637_average_of_levels_in_binary_tree

// #Easy #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree
// #Top_Interview_150_Binary_Tree_BFS #2025_05_24_Time_0_ms_(100.00%)_Space_8.58_MB_(36.34%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	levelSums := make(map[int][]float64)
	helper(root, levelSums, 0)
	result := make([]float64, len(levelSums))
	for level, pair := range levelSums {
		result[level] = pair[1] / pair[0]
	}
	return result
}

func helper(root *TreeNode, levelSums map[int][]float64, level int) {
	if root == nil {
		return
	}
	if _, exists := levelSums[level]; !exists {
		levelSums[level] = []float64{0, 0}
	}
	pair := levelSums[level]
	pair[0]++                    // count
	pair[1] += float64(root.Val) // sum
	levelSums[level] = pair
	helper(root.Left, levelSums, level+1)
	helper(root.Right, levelSums, level+1)
}
