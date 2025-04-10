package s0114_flatten_binary_tree_to_linked_list

// #Medium #Top_100_Liked_Questions #Depth_First_Search #Tree #Binary_Tree #Stack #Linked_List
// #Udemy_Linked_List #Top_Interview_150_Binary_Tree_General #Big_O_Time_O(N)_Space_O(N)
// #2024_03_19_Time_0_ms_(100.00%)_Space_3.1_MB_(5.75%)

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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	if root.Left != nil {
		rootRight := root.Right
		root.Right = root.Left
		t := root.Left
		for t.Right != nil {
			t = t.Right
		}
		t.Right = rootRight
	}
	root.Left = nil
}
