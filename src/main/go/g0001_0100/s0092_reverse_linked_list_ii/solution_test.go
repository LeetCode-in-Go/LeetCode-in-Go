package s0092_reverse_linked_list_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	result := reverseBetween(head, 2, 4)
	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 4, result.Next.Val)
	assert.Equal(t, 3, result.Next.Next.Val)
	assert.Equal(t, 2, result.Next.Next.Next.Val)
	assert.Equal(t, 5, result.Next.Next.Next.Next.Val)
}

func TestReverseBetween2(t *testing.T) {
	head := &ListNode{Val: 5}

	result := reverseBetween(head, 1, 1)
	assert.Equal(t, 5, result.Val)
}
