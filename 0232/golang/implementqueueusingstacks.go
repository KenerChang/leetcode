package implementqueueusingstacks

type MyQueue struct {
	stack []int
	top   int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append([]int{x}, this.stack...)
	if len(this.stack) == 1 {
		this.top = x
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	stackSize := len(this.stack)

	reverse := []int{}
	for i := 0; i < stackSize-1; i++ {
		reverse = append([]int{this.stack[i]}, reverse...)
	}

	popped := this.stack[stackSize-1]

	stack := []int{}
	for i := 0; i < len(reverse); i++ {
		stack = append([]int{reverse[i]}, stack...)
	}

	this.stack = stack
	if len(this.stack) > 0 {
		this.top = this.stack[len(this.stack)-1]
	}
	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.top
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
