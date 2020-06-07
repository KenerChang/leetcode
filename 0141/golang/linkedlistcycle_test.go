package linkedlistcycle

import (
	"testing"
)

func TestHas(t *testing.T) {
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
	target := true

	result := hasCycle(head)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestHasII(t *testing.T) {
	second := &ListNode{
		Val: 2,
	}

	head := &ListNode{
		Val:  1,
		Next: second,
	}
	second.Next = head
	target := true

	result := hasCycle(head)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestHasIII(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	target := false

	result := hasCycle(head)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestHasIV(t *testing.T) {
	target := false

	result := hasCycle(nil)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
