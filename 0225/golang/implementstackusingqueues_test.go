package implementstackusingqueues

import "testing"

func TestStackPush(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	target := 5
	if stack.queue[0] != target {
		t.Errorf("expect %d, got %d", target, stack.queue[0])
	}
}

func TestStackTop(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	target := 5
	if stack.Top() != target {
		t.Errorf("expect %d, got %d", target, stack.queue[0])
	}
}

func TestStackPop(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	result := stack.Pop()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestStackTopII(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	stack.Push(10)
	stack.Pop()
	result := stack.Top()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestStackEmpty(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	result := stack.Empty()

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}

	stack.Pop()
	result = stack.Empty()

	target = true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
func TestStack(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)
	result := stack.Top()

	target := 2
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}

	stack.Pop()
	empty := stack.Empty()

	emptyTarget := false
	if empty != emptyTarget {
		t.Errorf("expect %t, got %t", empty, emptyTarget)
	}
}
