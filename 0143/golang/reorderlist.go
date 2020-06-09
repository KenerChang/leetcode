package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slowP := head
	fastP := head

	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	// reverse second part
	middle := slowP
	var prev *ListNode
	for middle != nil {
		next := middle.Next

		middle.Next, prev = prev, middle
		middle = next
	}

	curP := head
	last := prev

	for last != nil {
		inorderNext, reverseNext := curP.Next, last.Next

		curP.Next, last.Next = last, inorderNext

		curP, last = inorderNext, reverseNext
	}
	slowP.Next = nil
}
