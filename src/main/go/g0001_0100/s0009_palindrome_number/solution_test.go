package s0009_palindrome_number

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome1(t *testing.T) {
	assert.Equal(t, true, isPalindrome(121))
}

func TestIsPalindrome2(t *testing.T) {
	assert.Equal(t, false, isPalindrome(-121))
}

func TestIsPalindrome3(t *testing.T) {
	assert.Equal(t, false, isPalindrome(10))
}
