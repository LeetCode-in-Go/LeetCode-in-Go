package s0106_construct_binary_tree_from_inorder_and_postorder_traversal

// #Medium #Array #Hash_Table #Tree #Binary_Tree #Divide_and_Conquer
// #Top_Interview_150_Binary_Tree_General #2025_05_17_Time_0_ms_(100.00%)_Space_6.06_MB_(38.37%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inIndex := []int{len(inorder) - 1}
	postIndex := []int{len(postorder) - 1}
	return helper(inorder, postorder, inIndex, postIndex, 1<<31-1)
}

func helper(in, post []int, index, pIndex []int, target int) *TreeNode {
	if index[0] < 0 || in[index[0]] == target {
		return nil
	}
	root := &TreeNode{Val: post[pIndex[0]]}
	pIndex[0]--
	root.Right = helper(in, post, index, pIndex, root.Val)
	index[0]--
	root.Left = helper(in, post, index, pIndex, target)
	return root
}
