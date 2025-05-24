package s0373_find_k_pairs_with_smallest_sums

// #Medium #Array #Heap_Priority_Queue #Top_Interview_150_Heap
// #2025_05_24_Time_21_ms_(90.40%)_Space_10.97_MB_(85.88%)

import "container/heap"

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add initial pairs
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(pq, &Node{
			sum:   int64(nums1[i]) + int64(nums2[0]),
			pair:  []int{nums1[i], nums2[0]},
			index: 0,
		})
	}

	result := make([][]int, 0, k)
	for i := 0; i < k && pq.Len() > 0; i++ {
		cur := heap.Pop(pq).(*Node)
		result = append(result, cur.pair)
		if cur.index+1 < len(nums2) {
			heap.Push(pq, &Node{
				sum:   int64(cur.pair[0]) + int64(nums2[cur.index+1]),
				pair:  []int{cur.pair[0], nums2[cur.index+1]},
				index: cur.index + 1,
			})
		}
	}

	return result
}
