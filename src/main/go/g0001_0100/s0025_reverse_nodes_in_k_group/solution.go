package s0025_reverse_nodes_in_k_group

// #Hard #Top_100_Liked_Questions #Linked_List #Recursion #Data_Structure_II_Day_13_Linked_List
// #Udemy_Linked_List #Top_Interview_150_Linked_List #Big_O_Time_O(n)_Space_O(k)
// #2025_04_27_Time_0_ms_(100.00%)_Space_5.54_MB_(20.53%)

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	j := 0
	len := head
	// Loop to check the length of the linked list. If the list is less than k, return as it is.
	for j < k {
		if len == nil {
			return head
		}
		len = len.Next
		j++
	}
	// Apply reverse linked list logic.
	c := head
	var n *ListNode
	var prev *ListNode
	i := 0
	// Traverse the while loop for k times to reverse the nodes in k groups.
	for i != k {
		n = c.Next
		c.Next = prev
		prev = c
		c = n
		i++
	}
	// head points to the first node of the k group, which now points to the next k group.
	// Recurse for the remaining linked list.
	head.Next = reverseKGroup(n, k)
	return prev
}
