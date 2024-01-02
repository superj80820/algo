// tags: stack, star1, medium

type MinStack struct {
	top *StackNode
}

type StackNode struct {
	data int
	next *StackNode
	min  int
}

func Constructor() MinStack {
	return MinStack{}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *MinStack) Push(val int) {
	if this.top == nil {
		this.top = &StackNode{data: val, min: val}
	} else {
		this.top = &StackNode{data: val, min: min(val, this.top.min), next: this.top}
	}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *MinStack) Pop() {
	this.top = this.top.next
}

// time complexity: O(1)
// space complexity: O(1)
func (this *MinStack) Top() int {
	return this.top.data
}

// time complexity: O(1)
// space complexity: O(1)
func (this *MinStack) GetMin() int {
	return this.top.min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// type MinStack struct {
// 	stack    []int
// 	minStack []int
// }

// func Constructor() MinStack {
// 	return MinStack{}
// }

// func (this *MinStack) Push(val int) {
// 	if len(this.minStack) == 0 {
// 		this.minStack = append(this.minStack, val)
// 	} else {
// 		this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
// 	}
// 	this.stack = append(this.stack, val)
// }

// func (this *MinStack) Pop() {
// 	this.stack = this.stack[:len(this.stack)-1]
// 	this.minStack = this.minStack[:len(this.minStack)-1]
// }

// func (this *MinStack) Top() int {
// 	return this.stack[len(this.stack)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.minStack[len(this.minStack)-1]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
