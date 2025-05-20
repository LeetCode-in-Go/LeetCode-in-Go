package s0122_best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 7, maxProfit(prices))
}

func TestMaxProfit2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 4, maxProfit(prices))
}

func TestMaxProfit3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	assert.Equal(t, 0, maxProfit(prices))
}
