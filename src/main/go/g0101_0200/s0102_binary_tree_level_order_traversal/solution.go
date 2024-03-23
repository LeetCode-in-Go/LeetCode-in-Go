package s0102_binary_tree_level_order_traversal

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Breadth_First_Search #Tree
// #Binary_Tree #Data_Structure_I_Day_11_Tree #Level_1_Day_6_Tree #Udemy_Tree_Stack_Queue
// #Big_O_Time_O(N)_Space_O(N) #2024_03_19_Time_3_ms_(70.81%)_Space_3.8_MB_(31.96%)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := make([][]int, 0)
	for len(queue) > 0 {
		levelSize := len(queue)
		temp := make([]int, 0)
		for levelSize > 0 {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelSize--
		}
		ans = append(ans, temp)
	}
	return ans
}
