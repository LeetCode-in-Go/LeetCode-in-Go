package s0015_3sum

import (
	"testing"
)

const result = "Expected %v, but got %v"

func TestThreeSum(t *testing.T) {
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	actual := threeSum([]int{-1, 0, 1, 2, -1, -4})

	if !equalSlices(expected, actual) {
		t.Errorf(result, expected, actual)
	}
}

func TestThreeSum2(t *testing.T) {
	expected := [][]int{}
	actual := threeSum([]int{})

	if !equalSlices(expected, actual) {
		t.Errorf(result, expected, actual)
	}
}

func TestThreeSum3(t *testing.T) {
	expected := [][]int{}
	actual := threeSum([]int{0})

	if !equalSlices(expected, actual) {
		t.Errorf(result, expected, actual)
	}
}

func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalIntSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
