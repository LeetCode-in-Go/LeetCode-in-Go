package s0052_n_queens_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	assert.Equal(t, 2, totalNQueens(4))
}

func TestTotalNQueens2(t *testing.T) {
	assert.Equal(t, 1, totalNQueens(1))
}
