package s0049_group_anagrams

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, expected, result)
}

func TestGroupAnagrams2(t *testing.T) {
	expected := [][]string{{""}}
	actual := groupAnagrams([]string{""})
	assert.Equal(t, expected, result)
}

func TestGroupAnagrams3(t *testing.T) {
	expected := [][]string{{"a"}}
	actual := groupAnagrams([]string{"a"})
	assert.Equal(t, expected, result)
}
