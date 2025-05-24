package s0433_minimum_genetic_mutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMutation(t *testing.T) {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	assert.Equal(t, 1, minMutation(start, end, bank))
}

func TestMinMutation2(t *testing.T) {
	start := "AACCGGTT"
	end := "AAACGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	assert.Equal(t, 2, minMutation(start, end, bank))
}

func TestMinMutation3(t *testing.T) {
	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	assert.Equal(t, 3, minMutation(start, end, bank))
}

func TestMinMutation4(t *testing.T) {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{}
	assert.Equal(t, -1, minMutation(start, end, bank))
}
