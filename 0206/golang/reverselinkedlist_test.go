package reverselinkedlist

import (
	"testing"
)

func verify(target, result *ListNode) bool {
	if target == nil && result == nil {
		return true
	}

	for target != nil && result != nil {
		if target.Val != result.Val {
			return false
		}

		target = target.Next
		result = result.Next
	}

	if target != nil || result != nil {
		return false
	}
	return true
}

func toList(list []int) *ListNode {
	var head *ListNode
	var node *ListNode
	for _, num := range list {
		next := &ListNode{
			Val: num,
		}
		if head == nil {
			head = next
			node = next
			continue
		}

		node.Next = next
		node = next
	}
	return head
}

func TestReverseList(t *testing.T) {
	head := toList([]int{1, 2, 3, 4, 5})
	target := toList([]int{5, 4, 3, 2, 1})

	result := reverseList(head)
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target.String(), result.String())
	}
}

func TestReverseListII(t *testing.T) {
	head := toList([]int{1})
	target := toList([]int{1})

	result := reverseList(head)
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target.String(), result.String())
	}
}

func TestReverseListIII(t *testing.T) {
	head := toList([]int{})
	target := toList([]int{})

	result := reverseList(head)
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target.String(), result.String())
	}
}
