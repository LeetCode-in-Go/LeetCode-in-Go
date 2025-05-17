package s0100_same_tree

// #Easy #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree #Level_2_Day_15_Tree
// #Udemy_Tree_Stack_Queue #Top_Interview_150_Binary_Tree_General
// #2025_05_17_Time_0_ms_(100.00%)_Space_4.26_MB_(12.80%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	b1 := isSameTree(p.Left, q.Left)
	b2 := isSameTree(p.Right, q.Right)
	return p.Val == q.Val && b1 && b2
}
