package s0236_lowest_common_ancestor_of_a_binary_tree

// #Medium #Top_100_Liked_Questions #Depth_First_Search #Tree #Binary_Tree
// #LeetCode_75_Binary_Tree/DFS #Data_Structure_II_Day_18_Tree #Udemy_Tree_Stack_Queue
// #Top_Interview_150_Binary_Tree_General #Big_O_Time_O(n)_Space_O(n)
// #2025_05_09_Time_5_ms_(89.93%)_Space_9.14_MB_(38.67%)

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
