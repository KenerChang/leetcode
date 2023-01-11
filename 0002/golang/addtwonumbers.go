package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Ints() []int {
	ints := []int{}

	curr := n

	for curr != nil {
		ints = append(ints, curr.Val)
		curr = curr.Next
	}

	return ints

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	longer := longerList(l1, l2)
	shorter := shorterList(l1, l2)

	if longer == nil {
		longer = l1
		shorter = l2
	}

	root := longer

	var carry int
	var last *ListNode
	for longer != nil || shorter != nil {
		if shorter != nil {
			longer.Val = longer.Val + shorter.Val + carry
		} else {
			longer.Val = longer.Val + carry
		}
		carry = 0

		if longer.Val >= 10 {
			longer.Val -= 10
			carry = 1
		}

		last = longer
		longer = longer.Next

		if shorter != nil {
			shorter = shorter.Next
		}
	}

	if carry > 0 {
		last.Next = &ListNode{
			Val: 1,
		}
	}

	return root
}

func longerList(l1, l2 *ListNode) *ListNode {
	l1p := l1
	l2p := l2

	for l1p != nil && l2p != nil {
		l1p = l1p.Next
		l2p = l2p.Next
	}

	if l1p == nil && l2p == nil {
		return nil
	} else if l1p == nil {
		return l2
	} else {
		return l1
	}
}

func shorterList(l1, l2 *ListNode) *ListNode {
	l1p := l1
	l2p := l2

	for l1p != nil && l2p != nil {
		l1p = l1p.Next
		l2p = l2p.Next
	}

	if l1p == nil && l2p == nil {
		return nil
	} else if l1p == nil {
		return l1
	} else {
		return l2
	}
}
