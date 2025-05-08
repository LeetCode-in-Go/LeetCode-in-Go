package s0148_sort_list

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Sorting #Two_Pointers #Linked_List
// #Divide_and_Conquer #Merge_Sort #Level_2_Day_4_Linked_List #Top_Interview_150_Divide_and_Conquer
// #Big_O_Time_O(log(N))_Space_O(log(N)) #2025_05_08_Time_7_ms_(87.91%)_Space_8.75_MB_(91.94%)

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
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMid(head)

	left := sortList(head)
	right := sortList(mid)

	return merge(left, right)
}

func findMid(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func merge(left, right *ListNode) *ListNode {
	pointer := new(ListNode)
	res := pointer
	for left != nil && right != nil {
		if left.Val < right.Val {
			pointer.Next = left
			left = left.Next
		} else {
			pointer.Next = right
			right = right.Next
		}
		pointer = pointer.Next
	}
	if left != nil {
		pointer.Next = left
	}
	if right != nil {
		pointer.Next = right
	}
	return res.Next
}
