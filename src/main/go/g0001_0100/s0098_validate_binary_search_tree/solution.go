package s0098_validate_binary_search_tree

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search #Tree #Binary_Tree
// #Binary_Search_Tree #Data_Structure_I_Day_14_Tree #Level_1_Day_8_Binary_Search_Tree
// #Udemy_Tree_Stack_Queue #Top_Interview_150_Binary_Search_Tree #Big_O_Time_O(N)_Space_O(log(N))
// #2024_03_15_Time_6_ms_(55.36%)_Space_5.4_MB_(59.64%)

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
func isValidBST(root *TreeNode) bool {
	return checkBST(root, nil, nil)
}

func checkBST(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && *min >= root.Val {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}
	return checkBST(root.Left, min, &root.Val) &&
		checkBST(root.Right, &root.Val, max)

}
