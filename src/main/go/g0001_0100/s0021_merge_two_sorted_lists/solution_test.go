package s0021_merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists1(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	expected := []int{1, 1, 2, 3, 4, 4}
	result := mergeTwoLists(l1, l2)
	assert.Equal(t, expected, listToArray(result))
}

func TestMergeTwoLists2(t *testing.T) {
	expected := []int{0, 0}
	node1 := &ListNode{Val: 0, Next: nil}
	node2 := &ListNode{Val: 0, Next: nil}
	result := mergeTwoLists(node1, node2)
	assert.Equal(t, expected, listToArray(result))
}

func listToArray(head *ListNode) []int {
	var arr []int
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}
