package s0105_construct_binary_tree_from_preorder_and_inorder_traversal

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table #Tree #Binary_Tree
// #Divide_and_Conquer #Data_Structure_II_Day_15_Tree #Big_O_Time_O(N)_Space_O(N)
// #2024_03_19_Time_0_ms_(100.00%)_Space_4.4_MB_(16.83%)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int)
	for i, value := range inorder {
		inorderIndexMap[value] = i
	}
	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inRootIndex := inorderIndexMap[rootVal]

		leftTreeSize := inRootIndex - inStart
		root.Left = build(preStart+1, preStart+leftTreeSize, inStart, inRootIndex-1)
		root.Right = build(preStart+leftTreeSize+1, preEnd, inRootIndex+1, inEnd)
		return root
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
