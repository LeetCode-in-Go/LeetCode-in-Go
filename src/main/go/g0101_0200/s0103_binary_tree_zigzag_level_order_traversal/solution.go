package s0103_binary_tree_zigzag_level_order_traversal

// #Medium #Top_Interview_Questions #Breadth_First_Search #Tree #Binary_Tree
// #Data_Structure_II_Day_15_Tree #Udemy_Tree_Stack_Queue #Top_Interview_150_Binary_Tree_BFS
// #2025_05_17_Time_0_ms_(100.00%)_Space_4.84_MB_(19.22%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root, nil}
	zig := true
	var level []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for node != nil {
			if zig {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if len(queue) > 0 {
				node = queue[0]
				queue = queue[1:]
			} else {
				node = nil
			}
		}
		result = append(result, level)
		zig = !zig
		level = []int{}
		if len(queue) > 0 {
			queue = append(queue, nil)
		}
	}
	return result
}
