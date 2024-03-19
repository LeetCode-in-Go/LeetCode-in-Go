package s0121_best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(prices)
	assert.Equal(t, 5, maxProfit, "The max profit should be 5")
}

func TestMaxProfit2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	maxProfit := maxProfit(prices)
	assert.Equal(t, 0, maxProfit, "The max profit should be 0")
}
