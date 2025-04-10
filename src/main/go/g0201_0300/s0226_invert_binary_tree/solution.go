package s0226_invert_binary_tree

// #Easy #Top_100_Liked_Questions #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree
// #Data_Structure_I_Day_12_Tree #Level_2_Day_6_Tree #Udemy_Tree_Stack_Queue
// #Top_Interview_150_Binary_Tree_General #Big_O_Time_O(n)_Space_O(n)
// #2024_03_22_Time_0_ms_(100.00%)_Space_2.4_MB_(5.25%)

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(temp)
	return root
}
