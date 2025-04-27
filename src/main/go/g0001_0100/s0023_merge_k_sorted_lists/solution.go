package s0023_merge_k_sorted_lists

// #Hard #Top_100_Liked_Questions #Top_Interview_Questions #Heap_Priority_Queue #Linked_List
// #Divide_and_Conquer #Merge_Sort #Top_Interview_150_Divide_and_Conquer
// #Big_O_Time_O(k*n*log(k))_Space_O(log(k)) #2025_04_27_Time_0_ms_(100.00%)_Space_7.12_MB_(27.34%)

import "container/heap"

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
func mergeKLists(lists []*ListNode) *ListNode {
	var h minHeap
	for _, node := range lists {
		for node != nil {
			h = append(h, node)
			node = node.Next
		}
	}
	heap.Init(&h)

	if len(h) == 0 {
		return nil
	}

	res := heap.Pop(&h).(*ListNode)
	cur := res
	for h.Len() > 0 {
		cur.Next = heap.Pop(&h).(*ListNode)
		cur.Next.Next = nil
		cur = cur.Next
	}

	return res
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(*ListNode))
}

func (h *minHeap) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
