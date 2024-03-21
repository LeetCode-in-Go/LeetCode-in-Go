package s0148_sort_list

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	listNode1 := &ListNode{Val: 4}
	listNode1.Next = &ListNode{Val: 2}
	listNode1.Next.Next = &ListNode{Val: 1}
	listNode1.Next.Next.Next = &ListNode{Val: 3}
	assert.Equal(t, "1, 2, 3, 4", sortList(listNode1).ToString())
}

func TestSortList2(t *testing.T) {
	listNode1 := &ListNode{Val: -1}
	listNode1.Next = &ListNode{Val: 5}
	listNode1.Next.Next = &ListNode{Val: 3}
	listNode1.Next.Next.Next = &ListNode{Val: 4}
	listNode1.Next.Next.Next.Next = &ListNode{Val: 0}
	assert.Equal(t, "-1, 0, 3, 4, 5", sortList(listNode1).ToString())
}

func TestSortList3(t *testing.T) {
	assert.Nil(t, sortList(nil))
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
