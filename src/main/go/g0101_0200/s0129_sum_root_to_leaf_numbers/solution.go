package s0129_sum_root_to_leaf_numbers

// #Medium #Depth_First_Search #Tree #Binary_Tree #Top_Interview_150_Binary_Tree_General
// #2025_05_20_Time_0_ms_(100.00%)_Space_4.18_MB_(78.36%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Solution struct {
	sum int
}

func sumNumbers(root *TreeNode) int {
	s := &Solution{sum: 0}
	s.recurseSum(root, 0)
	return s.sum
}

func (s *Solution) recurseSum(node *TreeNode, curNum int) {
	if node.Left == nil && node.Right == nil {
		s.sum += 10*curNum + node.Val
	} else {
		if node.Left != nil {
			s.recurseSum(node.Left, 10*curNum+node.Val)
		}
		if node.Right != nil {
			s.recurseSum(node.Right, 10*curNum+node.Val)
		}
	}
}
