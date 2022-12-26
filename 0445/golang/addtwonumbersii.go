package addtwonumbersii

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToInts() []int {
	ints := []int{}

	n := l
	for n != nil {
		ints = append(ints, n.Val)

		n = n.Next
	}

	return ints
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// reverse l1 and l2
	l1, l1Len := reverse(l1)
	l2, l2Len := reverse(l2)

	// add shorter list to longer list
	var lp, sp *ListNode
	if l1Len >= l2Len {
		lp = l1
		sp = l2
	} else {
		lp = l2
		sp = l1
	}

	newHead := lp

	var carry int
	for sp != nil && lp != nil {
		lp.Val += sp.Val + carry
		carry = 0

		if lp.Val >= 10 {
			lp.Val -= 10
			carry = 1
		}

		lp = lp.Next
		sp = sp.Next
	}

	for lp != nil {
		lp.Val += carry
		carry = 0

		if lp.Val >= 10 {
			lp.Val -= 10
			carry = 1
		}

		lp = lp.Next
	}

	// reverse lp
	newHead, _ = reverse(newHead)

	// check carry
	if carry > 0 {
		newHead = &ListNode{
			Val:  1,
			Next: newHead,
		}
	}

	return newHead
}

func reverse(l *ListNode) (*ListNode, int) {
	if l == nil {
		return nil, 0
	}

	var prev *ListNode
	n := l
	count := 0

	for n != nil {
		next := n.Next
		n.Next = prev
		prev = n
		n = next

		count++
	}

	return prev, count
}
