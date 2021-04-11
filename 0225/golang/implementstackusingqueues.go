package implementstackusingqueues

type MyStack struct {
	queue []int
	top   int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	this.top = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	total := len(this.queue) - 1
	reverse := make([]int, total)
	for i := 0; i < total; i++ {
		reverse[i] = this.queue[total-1-i]
	}

	popped := this.queue[total]

	remainedLast := len(reverse) - 1
	queue := make([]int, len(reverse))
	for i := 0; i < len(reverse); i++ {
		queue[i] = reverse[remainedLast-i]
	}

	this.queue = queue

	if len(this.queue) > 0 {
		this.top = this.queue[remainedLast]
	} else {
		this.top = 0
	}
	return popped
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
