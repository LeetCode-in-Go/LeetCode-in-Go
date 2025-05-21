package s0188_best_time_to_buy_and_sell_stock_iv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 2, maxProfit(2, []int{2, 4, 1}))
}

func TestMaxProfit2(t *testing.T) {
	assert.Equal(t, 7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
