package s0082_remove_duplicates_from_sorted_list_ii

// #Medium #Two_Pointers #Linked_List #Data_Structure_II_Day_11_Linked_List
// #Algorithm_II_Day_3_Two_Pointers #Top_Interview_150_Linked_List
// #2025_05_14_Time_0_ms_(100.00%)_Space_4.75_MB_(70.34%)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	prev := dummy
	prev.Next = head
	curr := head.Next
	for curr != nil {
		flagFoundDuplicate := false
		for curr != nil && prev.Next.Val == curr.Val {
			flagFoundDuplicate = true
			curr = curr.Next
		}
		if flagFoundDuplicate {
			prev.Next = curr
			if curr != nil {
				curr = curr.Next
			}
		} else {
			prev = prev.Next
			prev.Next = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}
