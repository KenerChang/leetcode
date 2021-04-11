package implementqueueusingstacks

import "testing"

func TestQueuePush(t *testing.T) {
	queue := Constructor()

	queue.Push(5)
	queue.Push(10)
	result := queue.Peek()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestQueuePop(t *testing.T) {
	queue := Constructor()

	queue.Push(5)
	result := queue.Pop()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestQueuePopII(t *testing.T) {
	queue := Constructor()

	queue.Push(5)
	queue.Push(10)
	result := queue.Pop()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := Constructor()

	queue.Push(5)
	queue.Push(10)
	result := queue.Peek()

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestQueueEmpty(t *testing.T) {
	queue := Constructor()

	result := queue.Empty()

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestQueueEmptyII(t *testing.T) {
	queue := Constructor()

	queue.Push(5)
	queue.Push(10)
	result := queue.Empty()

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestQueue(t *testing.T) {
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	peekResult := queue.Peek()

	peedTarget := 1
	if peekResult != peedTarget {
		t.Errorf("expect %d, got %d", peekResult, peedTarget)
	}

	popResult := queue.Pop()

	popTarget := 1
	if popResult != popTarget {
		t.Errorf("expect %d, got %d", popResult, popTarget)
	}

	emptyResult := queue.Empty()

	emptyTarget := false
	if emptyResult != emptyTarget {
		t.Errorf("expect %t, got %t", emptyResult, emptyTarget)
	}
}
