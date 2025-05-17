package s0112_path_sum

// #Easy #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree #Data_Structure_I_Day_12_Tree
// #Top_Interview_150_Binary_Tree_General #2025_05_17_Time_0_ms_(100.00%)_Space_6.84_MB_(7.76%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
