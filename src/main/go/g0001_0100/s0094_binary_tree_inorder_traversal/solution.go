package s0094_binary_tree_inorder_traversal

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search #Tree #Binary_Tree
// #Stack #Data_Structure_I_Day_10_Tree #Udemy_Tree_Stack_Queue #Big_O_Time_O(n)_Space_O(n)
// #2024_03_15_Time_0_ms_(100.00%)_Space_2.3_MB_(82.78%)

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
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node != nil {
			if node.Left != nil {
				travel(node.Left)
			}
			result = append(result, node.Val)
			if node.Right != nil {
				travel(node.Right)
			}
		}
	}
	travel(root)
	return result
}
