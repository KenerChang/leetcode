package minstack

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{
		stack: []int{},
	}

	return minStack
}

func (this *MinStack) Push(x int) {
	if x < this.min || len(this.stack) == 0 {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	last := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]

	if this.min == last && len(this.stack) > 0 {
		this.min = this.stack[0]
		for _, num := range this.stack {
			if num < this.min {
				this.min = num
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
