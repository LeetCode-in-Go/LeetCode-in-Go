package s0221_maximal_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	assert.Equal(t, 4, maximalSquare(input))
}

func TestMaximalSquare2(t *testing.T) {
	input := [][]byte{{'0', '1'}, {'1', '0'}}
	assert.Equal(t, 1, maximalSquare(input))
}

func TestMaximalSquare3(t *testing.T) {
	input := [][]byte{{'0'}}
	assert.Equal(t, 0, maximalSquare(input))
}
