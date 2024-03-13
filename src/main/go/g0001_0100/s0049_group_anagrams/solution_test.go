package s0049_group_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	actual := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert.Equal(t, expected, actual)
}

func TestGroupAnagrams2(t *testing.T) {
	expected := [][]string{{""}}
	actual := groupAnagrams([]string{""})
	assert.Equal(t, expected, actual)
}

func TestGroupAnagrams3(t *testing.T) {
	expected := [][]string{{"a"}}
	actual := groupAnagrams([]string{"a"})
	assert.Equal(t, expected, actual)
}
