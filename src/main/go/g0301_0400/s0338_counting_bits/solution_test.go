package s0338_counting_bits

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBits(t *testing.T) {
	result := countBits(2)
	expected := []int{0, 1, 1}
	assert.Equal(t, expected, result)
}

func TestCountBits2(t *testing.T) {
	result := countBits(5)
	expected := []int{0, 1, 1, 2, 1, 2}
	assert.Equal(t, expected, result)
}
