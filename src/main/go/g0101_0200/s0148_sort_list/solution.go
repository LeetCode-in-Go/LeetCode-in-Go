package s0148_sort_list

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Sorting #Two_Pointers #Linked_List
// #Divide_and_Conquer #Merge_Sort #Level_2_Day_4_Linked_List #Top_Interview_150_Divide_and_Conquer
// #Big_O_Time_O(log(N))_Space_O(log(N)) #2024_03_21_Time_40_ms_(97.89%)_Space_7_MB_(83.99%)

import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	curr := head
	a := []*ListNode{}
	for curr != nil {
		a = append(a, curr)
		curr = curr.Next
	}
	slices.SortFunc(a, func(l1, l2 *ListNode) int {
		return l1.Val - l2.Val
	})
	newNode := &ListNode{}
	curr = newNode
	for _, node := range a {
		curr.Next = node
		curr = node
	}
	curr.Next = nil
	return newNode.Next
}
