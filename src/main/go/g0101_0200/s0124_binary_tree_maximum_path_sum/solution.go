package s0124_binary_tree_maximum_path_sum

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Dynamic_Programming #Depth_First_Search
// #Tree #Binary_Tree #Udemy_Tree_Stack_Queue #Big_O_Time_O(N)_Space_O(N)
// #2024_03_19_Time_10_ms_(89.68%)_Space_7.8_MB_(94.84%)

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
func deepest(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := deepest(root.Left, ans)
	r := deepest(root.Right, ans)
	if l+root.Val > *ans {
		*ans = l + root.Val
	}
	if r+root.Val > *ans {
		*ans = r + root.Val
	}
	if l+r+root.Val > *ans {
		*ans = l + r + root.Val
	}
	if root.Val > *ans {
		*ans = root.Val
	}
	if l > r && l > 0 {
		return root.Val + l
	}
	if r > 0 {
		return root.Val + r
	}
	return root.Val
}

func maxPathSum(root *TreeNode) int {
	ans := -100000000
	deepest(root, &ans)
	return ans
}
