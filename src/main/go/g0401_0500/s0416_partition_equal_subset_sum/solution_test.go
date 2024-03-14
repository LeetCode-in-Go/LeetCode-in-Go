package s0416_partition_equal_subset_sum

import (
	"reflect"
	"testing"
)

func TestCanPartition(t *testing.T) {
	result := canPartition([]int{1, 5, 11, 5})
	expected := true
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func TestCanPartition2(t *testing.T) {
	result := canPartition([]int{1, 2, 3, 5})
	expected := false
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}
