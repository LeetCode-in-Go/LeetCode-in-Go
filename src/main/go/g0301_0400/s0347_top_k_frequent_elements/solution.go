package s0347_top_k_frequent_elements

// #Medium #Top_100_Liked_Questions #Array #Hash_Table #Sorting #Heap_Priority_Queue #Counting
// #Divide_and_Conquer #Quickselect #Bucket_Sort #Data_Structure_II_Day_20_Heap_Priority_Queue
// #Big_O_Time_O(n*log(n))_Space_O(k) #2025_05_10_Time_0_ms_(100.00%)_Space_7.96_MB_(62.32%)

import "sort"

type kvPair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	var pairs []kvPair
	for k, v := range m {
		pairs = append(pairs, kvPair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	var sortedKeys []int
	for _, pair := range pairs {
		sortedKeys = append(sortedKeys, pair.Key)
	}

	return sortedKeys[:k]
}
