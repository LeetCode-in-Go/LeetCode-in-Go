package s0230_kth_smallest_element_in_a_bst

// #Medium #Top_100_Liked_Questions #Depth_First_Search #Tree #Binary_Tree #Binary_Search_Tree
// #Data_Structure_II_Day_17_Tree #Level_2_Day_9_Binary_Search_Tree
// #Top_Interview_150_Binary_Search_Tree #Big_O_Time_O(n)_Space_O(n)
// #2025_05_09_Time_0_ms_(100.00%)_Space_8.52_MB_(17.71%)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	count int
	value int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func (n *Node) findkth(node *TreeNode, k int) {
	if node == nil {
		return
	}
	n.findkth(node.Left, k)
	n.count += 1
	if n.count == k {
		n.value = node.Val
		return
	}
	n.findkth(node.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
	n := &Node{
		count: 0,
		value: -1,
	}
	n.findkth(root, k)
	return n.value
}
