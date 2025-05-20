package s0123_best_time_to_buy_and_sell_stock_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	assert.Equal(t, 6, maxProfit(prices))
}

func TestMaxProfit2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 4, maxProfit(prices))
}

func TestMaxProfit3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	assert.Equal(t, 0, maxProfit(prices))
}

func TestMaxProfit4(t *testing.T) {
	var prices []int
	assert.Equal(t, 0, maxProfit(prices))
}
