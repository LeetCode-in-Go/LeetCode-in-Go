package s0142_linked_list_cycle_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	listNode1 := &ListNode{Val: 3}
	listNode1.Next = &ListNode{Val: 2}
	listNode1.Next.Next = &ListNode{Val: 0}
	listNode1.Next.Next.Next = &ListNode{Val: -4}
	listNode1.Next.Next.Next.Next = listNode1.Next
	assert.Equal(t, detectCycle(listNode1), listNode1.Next)
}

func TestDetectCycle2(t *testing.T) {
	listNode1 := &ListNode{Val: 1}
	listNode1.Next = &ListNode{Val: 2}
	listNode1.Next.Next = listNode1
	assert.Equal(t, detectCycle(listNode1), listNode1)
}

func TestDetectCycle3(t *testing.T) {
	listNode1 := &ListNode{Val: 1}
	assert.Nil(t, detectCycle(listNode1))
}
