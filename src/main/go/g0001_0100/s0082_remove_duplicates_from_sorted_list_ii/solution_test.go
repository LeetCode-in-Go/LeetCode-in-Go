package s0082_remove_duplicates_from_sorted_list_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	result := deleteDuplicates(head)
	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 2, result.Next.Val)
	assert.Equal(t, 5, result.Next.Next.Val)
}

func TestDeleteDuplicates2(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}
	result := deleteDuplicates(head)
	assert.Equal(t, 2, result.Val)
	assert.Equal(t, 3, result.Next.Val)
}
