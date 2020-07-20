package removelinkedlistelements

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	p := l
	result := ""
	for p != nil {
		result = fmt.Sprintf("%s->%d", result, p.Val)
		p = p.Next
	}
	return result
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	var newHead *ListNode
	for head != nil {
		if head.Val == val {
			if prev != nil {
				prev.Next = head.Next
			}
		} else {
			if newHead == nil {
				newHead = head
			}
			prev = head
		}

		head = head.Next
	}

	return newHead
}
