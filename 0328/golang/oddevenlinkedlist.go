package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toArray() []int {
	vals := []int{}

	n := l
	for n != nil {
		vals = append(vals, n.Val)
		n = n.Next
	}

	return vals
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Use two list to store odd nodes and even nodes
	odd := head
	evenHead := head.Next
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
