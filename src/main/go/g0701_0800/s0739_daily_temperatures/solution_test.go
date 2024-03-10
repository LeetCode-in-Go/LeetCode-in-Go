package s0739_daily_temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures1(t *testing.T) {
	result := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, result)
}

func TestDailyTemperatures2(t *testing.T) {
	result := dailyTemperatures([]int{30, 40, 50, 60})
	assert.Equal(t, []int{1, 1, 1, 0}, result)
}

func TestDailyTemperatures3(t *testing.T) {
	result := dailyTemperatures([]int{30, 60, 90})
	assert.Equal(t, []int{1, 1, 0}, result)
}
