package s0142_linked_list_cycle_ii

// #Medium #Top_100_Liked_Questions #Hash_Table #Two_Pointers #Linked_List
// #Data_Structure_II_Day_10_Linked_List #Level_1_Day_4_Linked_List #Udemy_Linked_List
// #Big_O_Time_O(N)_Space_O(1) #2025_05_07_Time_0_ms_(100.00%)_Space_6.80_MB_(18.18%)

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
func detectCycle(head *ListNode) *ListNode {
    m := map[*ListNode]bool {}
    for head != nil {
        if m[head] == true {
            return head
        }
        m[head] = true
        head = head.Next
    }
    return nil
}
