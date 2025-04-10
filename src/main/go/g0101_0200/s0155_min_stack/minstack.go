package s0155_min_stack

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Stack #Design
// #Data_Structure_II_Day_14_Stack_Queue #Programming_Skills_II_Day_18 #Level_2_Day_16_Design
// #Udemy_Design #Top_Interview_150_Stack #Big_O_Time_O(1)_Space_O(N)
// #2024_03_22_Time_12_ms_(79.58%)_Space_7.4_MB_(21.27%)

import "math"

type element struct {
	Val int
	Min int
}

type MinStack struct {
	Stack []element
	Min   int
}

func Constructor() MinStack {
	return MinStack{[]element{}, math.MaxInt32}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, element{val, this.Min})
	this.Min = min(this.Min, val)
}

func (this *MinStack) Pop() {
	l := len(this.Stack)
	this.Min = this.Stack[l-1].Min
	this.Stack = this.Stack[:l-1]
}

func (this *MinStack) Top() int {
	l := len(this.Stack)
	return this.Stack[l-1].Val
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
