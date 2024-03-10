package s0024_swap_nodes_in_pairs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	head := createSinglyLinkedList([]int{1, 2, 3, 4})
	expected := createSinglyLinkedList([]int{2, 1, 4, 3})
	output := swapPairs(head)
	assert.Equal(t, expected, output)
}

func TestSwapPairs2(t *testing.T) {
	head := createSinglyLinkedList([]int{1})
	expected := createSinglyLinkedList([]int{1})
	output := swapPairs(head)
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
