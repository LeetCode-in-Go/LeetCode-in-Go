package s0104_maximum_depth_of_binary_tree

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Depth_First_Search #Breadth_First_Search
// #Tree #Binary_Tree #Data_Structure_I_Day_11_Tree
// #Programming_Skills_I_Day_10_Linked_List_and_Tree #Udemy_Tree_Stack_Queue
// #Top_Interview_150_Binary_Tree_General #Big_O_Time_O(N)_Space_O(H)
// #2024_03_19_Time_0_ms_(100.00%)_Space_4.6_MB_(19.81%)

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
func maxDepth(root *TreeNode) int {
	return traverseAndCount(root, 0)
}

func traverseAndCount(root *TreeNode, counter int) int {
	if root == nil {
		return counter
	}

	counter++
	cl := traverseAndCount(root.Left, counter)
	cr := traverseAndCount(root.Right, counter)

	if cl > cr {
		return cl
	} else {
		return cr
	}
}
