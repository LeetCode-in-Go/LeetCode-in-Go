package s0437_path_sum_iii

// #Medium #Depth_First_Search #Tree #Binary_Tree #Level_2_Day_7_Tree #Big_O_Time_O(n)_Space_O(n)
// #2024_03_12_Time_7_ms_(74.18%)_Space_5.3_MB_(19.02%)

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
func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int]int)
	mp[0] = 1

	return solve(root, 0, targetSum, mp)
}

func solve(root *TreeNode, currSum, target int, mp map[int]int) int {
	if root == nil {
		return 0
	}

	currSum += root.Val
	res := 0
	if val, ok := mp[currSum-target]; ok {
		res = val
	}
	if val, ok := mp[currSum]; ok {
		mp[currSum] = val + 1
	} else {
		mp[currSum] = 1
	}

	res += solve(root.Left, currSum, target, mp) + solve(root.Right, currSum, target, mp)
	if val, ok := mp[currSum]; ok {
		mp[currSum] = val - 1
	}

	return res
}
