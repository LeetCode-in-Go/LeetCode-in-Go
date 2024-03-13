package s0051_n_queens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	expected := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
	actual := solveNQueens(4)
	assert.Equal(t, expected, actual)
}

func TestSolveNQueens2(t *testing.T) {
	expected := [][]string{{"Q"}}
	actual := solveNQueens(1)
	assert.Equal(t, expected, actual)
}
