package s0173_binary_search_tree_iterator

// #Medium #Tree #Binary_Tree #Stack #Design #Binary_Search_Tree #Iterator
// #Data_Structure_II_Day_17_Tree #Programming_Skills_II_Day_16 #Level_2_Day_9_Binary_Search_Tree
// #Top_Interview_150_Binary_Tree_General #2025_03_09_Time_15_ms_(100.00%)_Space_48.60_MB_(18.83%)

type BSTIterator struct {
	node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{node: root}
}

func (this *BSTIterator) Next() int {
	res := -1
	for this.node != nil {
		if this.node.Left != nil {
			rightMost := this.node.Left
			for rightMost.Right != nil {
				rightMost = rightMost.Right
			}
			rightMost.Right = this.node
			temp := this.node.Left
			this.node.Left = nil
			this.node = temp
		} else {
			res = this.node.Val
			this.node = this.node.Right
			return res
		}
	}
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.node != nil
}
