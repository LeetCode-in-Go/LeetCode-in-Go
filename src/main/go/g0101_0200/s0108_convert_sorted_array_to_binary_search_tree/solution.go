package s0108_convert_sorted_array_to_binary_search_tree

// #Easy #Top_Interview_Questions #Array #Tree #Binary_Tree #Binary_Search_Tree #Divide_and_Conquer
// #Data_Structure_II_Day_15_Tree #Level_2_Day_9_Binary_Search_Tree #Udemy_Tree_Stack_Queue
// #Top_Interview_150_Divide_and_Conquer #2025_05_17_Time_0_ms_(100.00%)_Space_5.39_MB_(26.99%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(num []int) *TreeNode {
	return makeTree(num, 0, len(num)-1)
}

func makeTree(num []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	midNode := &TreeNode{Val: num[mid]}
	midNode.Left = makeTree(num, left, mid-1)
	midNode.Right = makeTree(num, mid+1, right)
	return midNode
}
