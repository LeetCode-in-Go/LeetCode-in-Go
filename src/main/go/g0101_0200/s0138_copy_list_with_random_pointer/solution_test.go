package s0138_copy_list_with_random_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strconv"
)

func TestCopyRandomList(t *testing.T) {
	node7 := &Node{Val: 7}
	node13 := &Node{Val: 13}
	node11 := &Node{Val: 11}
	node10 := &Node{Val: 10}
	node1 := &Node{Val: 1}
	node7.Next = node13
	node13.Next = node11
	node11.Next = node10
	node10.Next = node1
	node1.Next = nil
	node7.Random = nil
	node13.Random = node7
	node11.Random = node1
	node10.Random = node11
	node1.Random = node7
	assert.Equal(t, "[[7,null],[13,0],[11,4],[10,2],[1,0]]", toString(copyRandomList(node7)))
}

func TestCopyRandomList2(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Next = node2
	node1.Random = node1
	node2.Next = nil
	node2.Random = node2
	assert.Equal(t, "[[1,1],[2,1]]", toString(copyRandomList(node1)))
}

func TestCopyRandomList3(t *testing.T) {
	node31 := &Node{Val: 3}
	node32 := &Node{Val: 3}
	node33 := &Node{Val: 3}
	node31.Next = node32
	node31.Random = nil
	node32.Next = node33
	node32.Random = node31
	node33.Next = nil
	node33.Random = nil
	assert.Equal(t, "[[3,null],[3,0],[3,null]]", toString(copyRandomList(node31)))
}

func toString(n *Node) string {
	var result string
	result += "[" + strconv.Itoa(n.Val)
	if n.Random == nil {
		result += ",null"
	} else {
		result += "," + strconv.Itoa(n.Random.Val)
	}
	result += "]"

	curr := n.Next
	for curr != nil {
		result += ",[" + strconv.Itoa(curr.Val)
		if curr.Random == nil {
			result += ",null"
		} else {
			RandomIndex := 0
			curr2 := n
			for curr2.Next != nil && curr2 != curr.Random {
				RandomIndex++
				curr2 = curr2.Next
			}
			result += "," + strconv.Itoa(RandomIndex)
		}
		result += "]"
		curr = curr.Next
	}
	return "[" + result + "]"
}
