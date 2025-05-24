package s0222_count_complete_tree_nodes

// #Easy #Depth_First_Search #Tree #Binary_Search #Binary_Tree #Binary_Search_II_Day_10
// #Top_Interview_150_Binary_Tree_General #2025_05_24_Time_0_ms_(100.00%)_Space_9.18_MB_(74.02%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := leftHeight(root)
	rightHeight := rightHeight(root)
	// case 1: When Height(Left sub-tree) = Height(right sub-tree) 2^h - 1
	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}

func leftHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + leftHeight(root.Left)
}

func rightHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + rightHeight(root.Right)
}
