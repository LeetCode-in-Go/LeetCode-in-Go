package s0125_valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
}

func TestIsPalindrome2(t *testing.T) {
	assert.False(t, isPalindrome("race a car"))
}

func TestIsPalindrome3(t *testing.T) {
	assert.True(t, isPalindrome(" "))
}

func TestIsPalindrome4(t *testing.T) {
	assert.True(t, isPalindrome(".,"))
}
