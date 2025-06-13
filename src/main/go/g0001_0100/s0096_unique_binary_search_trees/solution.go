package s0096_unique_binary_search_trees

// #Medium #Dynamic_Programming #Math #Tree #Binary_Tree #Binary_Search_Tree
// #Dynamic_Programming_I_Day_11 #Big_O_Time_O(n)_Space_O(1)
// #2025_05_06_Time_0_ms_(100.00%)_Space_3.85_MB_(79.74%)

func numTrees(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 2*n - i
		result /= i + 1
	}
	result /= n + 1
	return result
}
