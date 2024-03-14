package s0494_target_sum

import (
	"reflect"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	result := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	expected := 5
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}

func TestFindTargetSumWays2(t *testing.T) {
	result := findTargetSumWays([]int{1}, 1)
	expected := 1
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}
