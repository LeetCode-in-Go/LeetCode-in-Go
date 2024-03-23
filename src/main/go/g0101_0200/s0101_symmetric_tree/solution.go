package s0101_symmetric_tree

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search #Breadth_First_Search
// #Tree #Binary_Tree #Data_Structure_I_Day_11_Tree #Level_2_Day_15_Tree
// #Big_O_Time_O(N)_Space_O(log(N)) #2024_03_15_Time_0_ms_(100.00%)_Space_3_MB_(8.89%)

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
func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	leftRes := helper(left.Left, right.Right)
	rightRes := helper(left.Right, right.Left)
	return leftRes && rightRes
}

func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}
