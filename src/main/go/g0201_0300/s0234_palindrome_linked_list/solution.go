package s0234_palindrome_linked_list

// #Easy #Top_100_Liked_Questions #Two_Pointers #Stack #Linked_List #Recursion
// #Level_2_Day_3_Linked_List #Udemy_Linked_List #Big_O_Time_O(n)_Space_O(1)
// #2025_05_09_Time_1_ms_(87.36%)_Space_10.64_MB_(70.30%)

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
func isPalindrome(head *ListNode) bool {
	// Calculate the length
	len := 0
	right := head
	for right != nil {
		right = right.Next
		len++
	}

	// Reverse the right half of the list
	len /= 2
	right = head
	for i := 0; i < len; i++ {
		right = right.Next
	}
	var prev *ListNode
	for right != nil {
		next := right.Next
		right.Next = prev
		prev = right
		right = next
	}

	// Compare left half and right half
	for i := 0; i < len; i++ {
		if prev != nil && head.Val == prev.Val {
			head = head.Next
			prev = prev.Next
		} else {
			return false
		}
	}
	return true
}
