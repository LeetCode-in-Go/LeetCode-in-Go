package s0206_reverse_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	headActual := &ListNode{Val: 1}
	headActual.Next = &ListNode{Val: 2}
	headActual.Next.Next = &ListNode{Val: 3}
	headActual.Next.Next.Next = &ListNode{Val: 4}
	headActual.Next.Next.Next.Next = &ListNode{Val: 5}
	assert.Equal(t, reverseList(headActual).ToString(), "5, 4, 3, 2, 1")
}

func TestReverseList2(t *testing.T) {
	headActual := &ListNode{Val: 1}
	headActual.Next = &ListNode{Val: 2}
	assert.Equal(t, reverseList(headActual).ToString(), "2, 1")
}

func TestReverseList3(t *testing.T) {
	assert.Nil(t, reverseList(nil))
}

func (l *ListNode) ToString() string {
	if l == nil {
		return ""
	}
	result := fmt.Sprintf("%d", l.Val)
	for current := l.Next; current != nil; current = current.Next {
		result += fmt.Sprintf(", %d", current.Val)
	}
	return result
}
