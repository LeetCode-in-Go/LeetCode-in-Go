package s0086_partition_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	result := partition(head, 3)
	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 2, result.Next.Val)
	assert.Equal(t, 2, result.Next.Next.Val)
	assert.Equal(t, 4, result.Next.Next.Next.Val)
	assert.Equal(t, 3, result.Next.Next.Next.Next.Val)
	assert.Equal(t, 5, result.Next.Next.Next.Next.Next.Val)
}

func TestPartition2(t *testing.T) {
	head := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	result := partition(head, 2)
	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 2, result.Next.Val)
}
