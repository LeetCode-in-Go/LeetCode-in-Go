package s0295_find_median_from_data_stream

// #Hard #Top_100_Liked_Questions #Sorting #Two_Pointers #Design #Heap_Priority_Queue #Data_Stream
// #Top_Interview_150_Heap #Big_O_Time_O(n*log_n)_Space_O(n)
// #2025_05_10_Time_59_ms_(96.70%)_Space_25.46_MB_(10.38%)

func pushHeap(a *[]int, x int) {
	*a = append(*a, x)
	c := len(*a) - 1
	p := (c - 1) >> 1
	for 0 < c && (*a)[p] < (*a)[c] {
		(*a)[p], (*a)[c] = (*a)[c], (*a)[p]
		c = p
		p = (c - 1) >> 1
	}
}

func heapify(a []int, i int) {
	size := len(a)

	for i < size-1 {
		lc := min(size-1, (i<<1)+1)
		rc := min(size-1, (i<<1)+2)
		mc := lc
		if a[lc] < a[rc] {
			mc = rc
		}
		if a[mc] < a[i] {
			return
		}
		a[mc], a[i] = a[i], a[mc]
		i = mc
	}
}

func popHeap(a *[]int) int {
	size := len(*a)
	top := (*a)[0]
	(*a)[0], (*a)[size-1] = (*a)[size-1], (*a)[0]
	*a = (*a)[:size-1]
	heapify(*a, 0)
	return top
}

type MedianFinder struct {
	a []int
	b []int
}

func Constructor() MedianFinder {
	var obj MedianFinder
	return obj
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.a) == 0 || num < this.a[0] {
		pushHeap(&this.a, num)
		if 1+len(this.b) < len(this.a) {
			pushHeap(&this.b, -popHeap(&this.a))
		}
	} else {
		pushHeap(&this.b, -num)
		if 1+len(this.a) < len(this.b) {
			pushHeap(&this.a, -popHeap(&this.b))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.a) == len(this.b) {
		return (float64(this.a[0]) - float64(this.b[0])) * 0.5
	} else if len(this.a) < len(this.b) {
		return -float64(this.b[0])
	} else {
		return float64(this.a[0])
	}
}
