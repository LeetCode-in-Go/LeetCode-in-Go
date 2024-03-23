package s0234_palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	headActual := &ListNode{Val: 1}
	headActual.Next = &ListNode{Val: 2}
	headActual.Next.Next = &ListNode{Val: 2}
	headActual.Next.Next.Next = &ListNode{Val: 1}
	assert.True(t, isPalindrome(headActual))
}

func TestIsPalindrome2(t *testing.T) {
	headActual := &ListNode{Val: 1}
	headActual.Next = &ListNode{Val: 2}
	assert.False(t, isPalindrome(headActual))
}
