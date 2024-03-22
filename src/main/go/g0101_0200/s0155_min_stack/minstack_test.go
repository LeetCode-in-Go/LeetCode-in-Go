package s0155_min_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.GetMin(), "getMin should return -3")
	minStack.Pop()
	assert.Equal(t, 0, minStack.Top(), "top should return 0")
	assert.Equal(t, -2, minStack.GetMin(), "getMin should return -2")
}
