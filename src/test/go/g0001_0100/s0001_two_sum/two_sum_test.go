package s0001_two_sum

import (
	"reflect"
	"src/main/go/g0001_0100/s0001_two_sum"
	"testing"
)

var (
	case1, target1, result1 = []int{2, 7, 11, 15}, 9, []int{0, 1}
	case2, target2, result2 = []int{7, 11, 15}, 22, []int{0, 2}
	case3, target3, result3 = []int{11, 15}, 26, []int{0, 1}
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(twoSum(case1, target1), result1) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum(case2, target2), result2) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum(case3, target3), result3) {
		t.Fail()
	}
	if twoSum(nil, 0) != nil {
		t.Fail()
	}
}