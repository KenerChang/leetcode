package sortlist

// import (
// 	"fmt"
// )

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slowP := head
	fastP := head.Next
	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	next := slowP.Next
	slowP.Next = nil

	// fmt.Printf("next: %d\n", next.Val)

	head1 := sortList(head)
	head2 := sortList(next)

	var newHead, tail, node *ListNode
	for head1 != nil || head2 != nil {
		if head1 == nil {
			node = head2
			head2 = head2.Next
		} else if head2 == nil {
			node = head1
			head1 = head1.Next
		} else if head1.Val >= head2.Val {
			node = head2
			head2 = head2.Next
		} else {
			node = head1
			head1 = head1.Next
		}

		if newHead == nil {
			newHead = node
			tail = newHead
		} else {
			tail.Next = node
			tail = node
		}
	}

	return newHead
}
