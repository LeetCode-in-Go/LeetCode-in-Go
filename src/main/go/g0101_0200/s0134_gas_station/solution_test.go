package s0134_gas_station

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))
}

func TestCanCompleteCircuit2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	assert.Equal(t, -1, canCompleteCircuit(gas, cost))
}

func TestCanCompleteCircuit3(t *testing.T) {
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	assert.Equal(t, 4, canCompleteCircuit(gas, cost))
}

func TestCanCompleteCircuit4(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 0, canCompleteCircuit(gas, cost))
}
