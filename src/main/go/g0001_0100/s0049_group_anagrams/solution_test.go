package s0049_group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	actual := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	sort.Slice(actual, func(i, j int) bool {
		return len(actual[i]) < len(actual[j])
	})
	for i := range actual {
		sort.Strings(actual[i])
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func TestGroupAnagrams2(t *testing.T) {
	expected := [][]string{{""}}
	actual := groupAnagrams([]string{""})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func TestGroupAnagrams3(t *testing.T) {
	expected := [][]string{{"a"}}
	actual := groupAnagrams([]string{"a"})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
