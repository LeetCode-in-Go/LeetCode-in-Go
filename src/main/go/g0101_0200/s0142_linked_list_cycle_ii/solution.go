package s0142_linked_list_cycle_ii

// #Medium #Top_100_Liked_Questions #Hash_Table #Two_Pointers #Linked_List
// #Data_Structure_II_Day_10_Linked_List #Level_1_Day_4_Linked_List #Udemy_Linked_List
// #Big_O_Time_O(N)_Space_O(1) #2024_03_20_Time_5_ms_(69.02%)_Space_3.8_MB_(98.54%)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// intersected inside the loop.
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
