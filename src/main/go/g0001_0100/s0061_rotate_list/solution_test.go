package s0061_rotate_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := rotateRight(head, 2)
	assert.Equal(t, 4, result.Val)
	assert.Equal(t, 5, result.Next.Val)
	assert.Equal(t, 1, result.Next.Next.Val)
	assert.Equal(t, 2, result.Next.Next.Next.Val)
	assert.Equal(t, 3, result.Next.Next.Next.Next.Val)
}

func TestRotateRight2(t *testing.T) {
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	result := rotateRight(head, 4)
	assert.Equal(t, 2, result.Val)
	assert.Equal(t, 0, result.Next.Val)
	assert.Equal(t, 1, result.Next.Next.Val)
}
