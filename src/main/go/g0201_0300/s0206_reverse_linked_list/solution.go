package s0206_reverse_linked_list

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Linked_List #Recursion
// #LeetCode_75_LinkedList #Data_Structure_I_Day_8_Linked_List
// #Algorithm_I_Day_10_Recursion_Backtracking #Level_1_Day_3_Linked_List #Udemy_Linked_List
// #Big_O_Time_O(N)_Space_O(1) #2025_05_09_Time_0_ms_(100.00%)_Space_4.36_MB_(95.25%)

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
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
