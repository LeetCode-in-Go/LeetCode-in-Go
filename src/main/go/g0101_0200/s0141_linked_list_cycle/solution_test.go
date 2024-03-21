package s0141_linked_list_cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	listNode1 := &ListNode{Val: 3}
	listNode1.Next = &ListNode{Val: 2}
	listNode1.Next.Next = &ListNode{Val: 0}
	listNode1.Next.Next.Next = &ListNode{Val: -4}
	listNode1.Next.Next.Next.Next = listNode1.Next
	assert.Equal(t, true, hasCycle(listNode1))
}

func TestHasCycle2(t *testing.T) {
	listNode1 := &ListNode{Val: 1}
	listNode1.Next = &ListNode{Val: 2}
	listNode1.Next.Next = listNode1
	assert.Equal(t, true, hasCycle(listNode1))
}

func TestHasCycle3(t *testing.T) {
	listNode1 := &ListNode{Val: 1}
	assert.Equal(t, false, hasCycle(listNode1))
}
