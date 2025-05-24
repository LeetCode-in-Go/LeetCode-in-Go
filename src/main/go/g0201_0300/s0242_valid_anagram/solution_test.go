package s0242_valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
}

func TestIsAnagram2(t *testing.T) {
	assert.False(t, isAnagram("rat", "car"))
}

func TestIsAnagram3(t *testing.T) {
	assert.False(t, isAnagram("a", "ab"))
}
