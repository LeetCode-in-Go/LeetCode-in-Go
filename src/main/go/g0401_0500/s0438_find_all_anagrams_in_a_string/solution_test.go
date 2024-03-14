package s0438_find_all_anagrams_in_a_string

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	result := Solution{}.findAnagrams("cbaebabacd", "abc")
	expected := []int{0, 6}
	sort.Ints(result)
	sort.Ints(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func TestFindAnagrams2(t *testing.T) {
	result := Solution{}.findAnagrams("abab", "ab")
	expected := []int{0, 1, 2}
	sort.Ints(result)
	sort.Ints(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}
