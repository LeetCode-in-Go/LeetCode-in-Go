package s0086_partition_list

// #Medium #Two_Pointers #Linked_List #Top_Interview_150_Linked_List
// #2025_05_17_Time_0_ms_(100.00%)_Space_4.30_MB_(62.68%)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	nHead := &ListNode{Val: 0}
	nTail := &ListNode{Val: 0}
	ptr := nTail
	temp := nHead
	for head != nil {
		nNext := head.Next
		if head.Val < x {
			nHead.Next = head
			nHead = nHead.Next
		} else {
			nTail.Next = head
			nTail = nTail.Next
		}
		head = nNext
	}
	nTail.Next = nil
	nHead.Next = ptr.Next
	return temp.Next
}
