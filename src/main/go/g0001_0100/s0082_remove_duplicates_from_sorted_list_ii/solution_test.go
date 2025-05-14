package s0082_remove_duplicates_from_sorted_list_ii

import (
	"testing"

	"github.com/LeetCode-in-Go/LeetCode-in-Go/src/main/go/com_github_leetcode/list_node"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &list_node.ListNode{Val: 1, Next: &list_node.ListNode{Val: 2, Next: &list_node.ListNode{Val: 3, Next: &list_node.ListNode{Val: 3, Next: &list_node.ListNode{Val: 4, Next: &list_node.ListNode{Val: 4, Next: &list_node.ListNode{Val: 5}}}}}}}
	result := deleteDuplicates(head)
	assert.Equal(t, 1, result.Val)
	assert.Equal(t, 2, result.Next.Val)
	assert.Equal(t, 5, result.Next.Next.Val)
}

func TestDeleteDuplicates2(t *testing.T) {
	head := &list_node.ListNode{Val: 1, Next: &list_node.ListNode{Val: 1, Next: &list_node.ListNode{Val: 1, Next: &list_node.ListNode{Val: 2, Next: &list_node.ListNode{Val: 3}}}}}
	result := deleteDuplicates(head)
	assert.Equal(t, 2, result.Val)
	assert.Equal(t, 3, result.Next.Val)
}
