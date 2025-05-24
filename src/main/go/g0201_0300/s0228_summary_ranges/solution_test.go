package s0228_summary_ranges

import (
	"github.com/stretchr/testify/assert"
    "testing"
)

func TestSummaryRanges(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	expected := []string{"0->2", "4->5", "7"}
	assert.Equal(t, expected, summaryRanges(nums))
}

func TestSummaryRanges2(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	expected := []string{"0", "2->4", "6", "8->9"}
	assert.Equal(t, expected, summaryRanges(nums))
}

func TestSummaryRanges3(t *testing.T) {
	nums := []int{}
	expected := []string{}
	assert.Equal(t, expected, summaryRanges(nums))
}
