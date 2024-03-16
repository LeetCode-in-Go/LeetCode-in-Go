package s0322_coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	result := coinChange([]int{1, 2, 5}, 11)
	expected := 3
	assert.Equal(t, expected, result)
}

func TestCoinChange2(t *testing.T) {
	result := coinChange([]int{2}, 3)
	expected := -1
	assert.Equal(t, expected, result)
}

func TestCoinChange3(t *testing.T) {
	result := coinChange([]int{1}, 0)
	expected := 0
	assert.Equal(t, expected, result)
}
