package s0274_h_index

// #Medium #Array #Sorting #Counting_Sort #Top_Interview_150_Array/String
// #2025_05_24_Time_0_ms_(100.00%)_Space_4.12_MB_(82.63%)

func hIndex(citations []int) int {
	len := len(citations)
	freqArray := make([]int, len+1)
	for _, citation := range citations {
		if citation > len {
			freqArray[len]++
		} else {
			freqArray[citation]++
		}
	}
	totalSoFar := 0
	for k := len; k >= 0; k-- {
		totalSoFar += freqArray[k]
		if totalSoFar >= k {
			return k
		}
	}
	return -1
}
