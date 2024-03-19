package s0295_find_median_from_data_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	assert.Equal(t, 1.5, medianFinder.FindMedian())

	medianFinder.AddNum(3)
	assert.Equal(t, 2.0, medianFinder.FindMedian())
}

func TestMedianFinder2(t *testing.T) {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(3)
	medianFinder.AddNum(-1)
	assert.Equal(t, 1.0, medianFinder.FindMedian())
}
