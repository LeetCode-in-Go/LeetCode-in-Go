package s0061_rotate_list

// #Medium #Two_Pointers #Linked_List #Programming_Skills_II_Day_16 #Udemy_Linked_List
// #Top_Interview_150_Linked_List #2025_03_04_Time_0_ms_(100.00%)_Space_42.42_MB_(78.37%)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	tail := head
	count := 1
	for tail != nil && tail.Next != nil {
		count++
		tail = tail.Next
	}
	times := k % count
	if times == 0 {
		return head
	}
	temp := head
	for i := 1; i <= count-times-1 && temp != nil; i++ {
		temp = temp.Next
	}
	var newHead *ListNode
	if temp != nil && tail != nil {
		newHead = temp.Next
		temp.Next = nil
		tail.Next = head
	}
	return newHead
}
