package s0160_intersection_of_two_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	intersectionListNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	nodeA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectionListNode}}
	nodeB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersectionListNode}}}
	assert.Equal(t, 8, getIntersectionNode(nodeA, nodeB).Val)
}

func TestGetIntersectionNode2(t *testing.T) {
	nodeA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	nodeB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}
	assert.Nil(t, getIntersectionNode(nodeA, nodeB))
}
