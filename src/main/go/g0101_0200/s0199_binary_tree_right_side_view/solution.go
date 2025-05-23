package s0199_binary_tree_right_side_view

// #Medium #Top_100_Liked_Questions #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree
// #LeetCode_75_Binary_Tree/BFS #Data_Structure_II_Day_16_Tree #Level_2_Day_15_Tree
// #Top_Interview_150_Binary_Tree_BFS #2025_05_22_Time_0_ms_(100.00%)_Space_4.29_MB_(85.87%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var list []int
	recurse(root, 0, &list)
	return list
}

func recurse(node *TreeNode, level int, list *[]int) {
	if node != nil {
		if len(*list) < level+1 {
			*list = append(*list, node.Val)
		}
		recurse(node.Right, level+1, list)
		recurse(node.Left, level+1, list)
	}
}
