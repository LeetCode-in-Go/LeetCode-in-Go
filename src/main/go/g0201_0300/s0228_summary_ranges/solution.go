package s0228_summary_ranges

// #Easy #Array #Top_Interview_150_Intervals #2025_05_24_Time_0_ms_(100.00%)_Space_3.94_MB_(33.05%)

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	ranges := make([]string, 0)
	if len(nums) == 0 {
		return ranges
	}
	// size of array
	n := len(nums)
	// start of range
	a := nums[0]
	// end of range
	b := a
	for i := 1; i < n; i++ {
		// we need to make a decision if the next element
		// will expand the range
		// i starts at 1, not 0, because 1 is the next
		// candidate for expanding the range
		if nums[i] != b+1 {
			// only when our next element does not expand the range
			// do we add the range a->b to our list of ranges
			if a == b {
				ranges = append(ranges, fmt.Sprintf("%d", a))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d->%d", a, b))
			}
			// since nums[i] is not accounted for by our range a->b
			// because nums[i] is not b+1, we need to set a and b
			// to this new range start point of bigger than b+1
			// maybe it is b+2? b+3? b+4? all we know is it is not b+1
			a = nums[i]
			b = a
		} else {
			// if the next element expands our range we do so
			b++
		}
	}
	// the only range that is not accounted for at this point is the last range
	// if our a and b are not equal then we add the range accordingly
	if a == b {
		ranges = append(ranges, fmt.Sprintf("%d", a))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d->%d", a, b))
	}
	return ranges
}
