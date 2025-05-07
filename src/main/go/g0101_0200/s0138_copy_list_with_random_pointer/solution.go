package s0138_copy_list_with_random_pointer

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Hash_Table #Linked_List
// #Programming_Skills_II_Day_14 #Udemy_Linked_List #Top_Interview_150_Linked_List
// #Big_O_Time_O(N)_Space_O(N) #2025_05_07_Time_0_ms_(100.00%)_Space_5.44_MB_(28.63%)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := map[*Node]*Node{}

	n := head
	for n != nil {
		m[n] = &Node{Val: n.Val}
		n = n.Next
	}

	n = head
	for n != nil {
		m[n].Next = m[n.Next]
		m[n].Random = m[n.Random]
		n = n.Next
	}

	return m[head]
}
