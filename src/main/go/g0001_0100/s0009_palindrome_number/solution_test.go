package s0009_palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
}

func TestIsPalindrome2(t *testing.T) {
	assert.False(t, isPalindrome(-121))
}

func TestIsPalindrome3(t *testing.T) {
	assert.False(t, isPalindrome(10))
}
