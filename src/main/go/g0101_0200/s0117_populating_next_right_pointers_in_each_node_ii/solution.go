package s0117_populating_next_right_pointers_in_each_node_ii

// #Medium #Depth_First_Search #Breadth_First_Search #Tree #Binary_Tree #Linked_List
// #Algorithm_II_Day_7_Breadth_First_Search_Depth_First_Search
// #Top_Interview_150_Binary_Tree_General #2025_05_17_Time_5_ms_(48.45%)_Space_8.12_MB_(54.08%)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else if root.Next != nil {
			root.Left.Next = adjacentRightNode(root.Next)
		}
	}
	if root.Right != nil {
		root.Right.Next = adjacentRightNode(root.Next)
	}
	connect(root.Right)
	connect(root.Left)
	return root
}

func adjacentRightNode(root *Node) *Node {
	temp := root
	for temp != nil {
		if temp.Left == nil && temp.Right == nil {
			temp = temp.Next
		} else {
			if temp.Left != nil {
				return temp.Left
			}
			if temp.Right != nil {
				return temp.Right
			}
		}
	}
	return nil
}
