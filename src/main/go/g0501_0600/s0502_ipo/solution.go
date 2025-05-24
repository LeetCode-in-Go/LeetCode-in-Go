package s0502_ipo

// #Hard #Array #Sorting #Greedy #Heap_Priority_Queue #Top_Interview_150_Heap
// #2025_05_24_Time_62_ms_(91.95%)_Space_15.28_MB_(54.36%)

import (
	"container/heap"
)

type profitAndCapital struct {
	profit, capital int
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type minHeap []profitAndCapital

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(profitAndCapital))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, currentCapital int, profits []int, capital []int) int {
	n := len(profits)
	profitHeap := &maxHeap{}
	capitalHeap := &minHeap{}
	// we gonna to a maxHeap per capital needed
	// in the max heap only store the project that we can do with the current capital
	// and in the minHeap we store the project that we can't do with the current capital
	for i := 0; i < n; i++ {
		if capital[i] <= currentCapital {
			heap.Push(profitHeap, profits[i])
		} else {
			heap.Push(capitalHeap, profitAndCapital{capital: capital[i], profit: profits[i]})
		}
	}
	// now what we need is to find a project that will maximise our profit
	// so depending on the curCapital we have we need to find the projet with the biggest profit until with reached k
	for k > 0 {
		// find the project that we can do with the current capital
		if profitHeap.Len() == 0 {
			break
		}
		currentCapital += heap.Pop(profitHeap).(int)
		for capitalHeap.Len() > 0 && (*capitalHeap)[0].capital <= currentCapital {
			x := heap.Pop(capitalHeap).(profitAndCapital)
			heap.Push(profitHeap, x.profit)
		}
		k--
	}
	return currentCapital
}
