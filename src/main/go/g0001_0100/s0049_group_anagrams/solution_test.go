package s0049_group_anagrams

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	actual := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestGroupAnagrams2(t *testing.T) {
	expected := [][]string{{""}}
	actual := groupAnagrams([]string{""})
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestGroupAnagrams3(t *testing.T) {
	expected := [][]string{{"a"}}
	actual := groupAnagrams([]string{"a"})
	assert.True(t, reflect.DeepEqual(expected, actual))
}
