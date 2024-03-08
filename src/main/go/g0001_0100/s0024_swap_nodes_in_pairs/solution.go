package s0024_swap_nodes_in_pairs

// #Medium #Top_100_Liked_Questions #Linked_List #Recursion #Data_Structure_II_Day_12_Linked_List
// #Udemy_Linked_List #Big_O_Time_O(n)_Space_O(1)
// #2024_03_08_Time_0_ms_(100.00%)_Space_2.1_MB_(95.56%)

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second, third := head, head.Next, swapPairs(head.Next.Next)
	first.Next = third
	second.Next = first
	return second
}
