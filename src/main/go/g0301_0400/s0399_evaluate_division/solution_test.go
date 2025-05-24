package s0399_evaluate_division

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	expected := []float64{6.0, 0.5, -1.0, 1.0, -1.0}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}

func TestCalcEquation2(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	expected := []float64{3.75, 0.4, 5.0, 0.2}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}

func TestCalcEquation3(t *testing.T) {
	equations := [][]string{{"a", "b"}}
	values := []float64{0.5}
	queries := [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}
	expected := []float64{0.5, 2.0, -1.0, -1.0}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}
