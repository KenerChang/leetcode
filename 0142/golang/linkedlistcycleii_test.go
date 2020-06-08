package linkedlistcycleii

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	second := &ListNode{
		Val: 0,
	}
	fourth := &ListNode{
		Val:  -4,
		Next: second,
	}
	third := &ListNode{
		Val:  0,
		Next: fourth,
	}
	second.Next = third

	head := &ListNode{
		Val:  3,
		Next: second,
	}
	target := second

	result := detectCycle(head)
	if target != result {
		t.Errorf("expect %v, got: %v", *target, *result)
	}
}

func TestDetectCycleII(t *testing.T) {
	second := &ListNode{
		Val: 2,
	}

	head := &ListNode{
		Val:  1,
		Next: second,
	}
	second.Next = head
	target := head

	result := detectCycle(head)
	if target != result {
		t.Errorf("expect %v, got: %v", *target, *result)
	}
}

func TestDetectCycleIII(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	var target *ListNode

	result := detectCycle(head)
	if target != result {
		t.Errorf("expect %p, got: %p", target, result)
	}
}

func TestDetectCycleIV(t *testing.T) {
	var target *ListNode

	result := detectCycle(nil)
	if target != result {
		t.Errorf("expect %v, got: %v", *target, *result)
	}
}
