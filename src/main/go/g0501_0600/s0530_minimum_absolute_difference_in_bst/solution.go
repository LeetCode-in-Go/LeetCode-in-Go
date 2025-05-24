package s0530_minimum_absolute_difference_in_bst

// #Easy #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree #Binary_Search_Tree
// #Top_Interview_150_Binary_Search_Tree #2025_05_24_Time_0_ms_(100.00%)_Space_8.46_MB_(71.12%)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	ans := 1<<31 - 1
	prev := 1<<31 - 1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != 1<<31-1 {
			diff := abs(root.Val - prev)
			if diff < ans {
				ans = diff
			}
		}
		prev = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
