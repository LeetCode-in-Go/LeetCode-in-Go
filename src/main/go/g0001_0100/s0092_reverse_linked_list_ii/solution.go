package s0092_reverse_linked_list_ii

// #Medium #Linked_List #Top_Interview_150_Linked_List
// #2025_05_17_Time_0_ms_(100.00%)_Space_4.09_MB_(56.36%)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var prev *ListNode
	temp := head
	var start *ListNode
	k := left
	for temp != nil && k > 1 {
		prev = temp
		temp = temp.Next
		k--
	}
	if left > 1 && prev != nil {
		prev.Next = nil
	}
	var prev1 *ListNode
	start = temp
	for temp != nil && right-left >= 0 {
		prev1 = temp
		temp = temp.Next
		right--
	}
	if prev1 != nil {
		prev1.Next = nil
	}
	if left > 1 && prev != nil {
		prev.Next = reverse(start)
	} else {
		head = reverse(start)
		prev = head
	}
	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = temp
	return head
}

func reverse(head *ListNode) *ListNode {
	var p, q, r *ListNode
	p = head
	for p != nil {
		q = p.Next
		p.Next = r
		r = p
		p = q
	}
	return r
}

func decode(s string) string {
	return s
}
