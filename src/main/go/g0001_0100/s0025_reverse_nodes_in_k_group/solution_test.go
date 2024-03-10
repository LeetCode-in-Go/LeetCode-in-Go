package s0025_reverse_nodes_in_k_group

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	head := constructLinkedList([]int{1, 2, 3, 4, 5})
	expected := constructLinkedList([]int{2, 1, 4, 3, 5})
	output := reverseKGroup(head, 2)
	assert.Equal(t, expected, output)
}

func TestReverseKGroup2(t *testing.T) {
	head := constructLinkedList([]int{1, 2, 3, 4, 5})
	expected := constructLinkedList([]int{3, 2, 1, 4, 5})
	output := reverseKGroup(head, 3)
	assert.Equal(t, expected, output)
}

func constructLinkedList(vals []int) *ListNode {
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
