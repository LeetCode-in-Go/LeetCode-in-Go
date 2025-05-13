package s0030_substring_with_concatenation_of_all_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	assert.Equal(t, []int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func TestFindSubstring2(t *testing.T) {
	assert.Equal(t, []int{}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
}

func TestFindSubstring3(t *testing.T) {
	assert.Equal(t, []int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
