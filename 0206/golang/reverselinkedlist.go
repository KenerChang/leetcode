package reverselinkedlist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	head := ln
	result := ""
	for head != nil {
		result = fmt.Sprintf("%s->%d", result, head.Val)
		head = head.Next
	}
	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		if prev != nil {
			cur.Next = prev
		} else {
			cur.Next = nil
		}
		prev = cur
		cur = next
	}
	return prev
}
