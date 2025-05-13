package s0028_find_the_index_of_the_first_occurrence_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
}

func TestStrStr2(t *testing.T) {
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
}
