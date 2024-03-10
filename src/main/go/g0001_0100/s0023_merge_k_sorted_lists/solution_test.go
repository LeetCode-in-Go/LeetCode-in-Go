package s0023_merge_k_sorted_lists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	head1 := createSinglyLinkedList([]int{1, 4, 5})
	head2 := createSinglyLinkedList([]int{1, 3, 4})
	head3 := createSinglyLinkedList([]int{2, 6})
	expected := createSinglyLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	output := mergeKLists([]*ListNode{head1, head2, head3})
	assert.Equal(t, expected, output)
}

func TestMergeKLists2(t *testing.T) {
	head1 := createSinglyLinkedList([]int{1, 3, 5, 7, 11})
	head2 := createSinglyLinkedList([]int{2, 8, 12})
	head3 := createSinglyLinkedList([]int{4, 6, 9, 10})
	expected := createSinglyLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	output := mergeKLists([]*ListNode{head1, head2, head3})
	assert.Equal(t, expected, output)
}

// ListNode definition should be included here

func createSinglyLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}
