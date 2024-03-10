package s0019_remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	expected := []int{1, 2, 3, 5}
	result := removeNthFromEnd(node1, 2)
	assert.Equal(t, expected, listToArray(result))
}

func TestRemoveNthFromEnd2(t *testing.T) {
	node1 := &ListNode{Val: 1}
	emptyResult := removeNthFromEnd(node1, 1)
	assert.Nil(t, emptyResult)
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
