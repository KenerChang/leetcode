package plusonelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToInts() []int {
	n := l

	ints := []int{}
	for n != nil {
		ints = append(ints, n.Val)
		n = n.Next
	}
	return ints
}

func plusOne(head *ListNode) *ListNode {
	// reverse the list
	if head == nil {
		return nil
	}

	newHead := reverse(head)

	// add one and handle propagation
	n := newHead
	for n != nil && n.Val+1 >= 10 {
		n.Val = 0
		n = n.Next
	}

	if n != nil {
		n.Val++
	}

	newHead = reverse(newHead)

	if n == nil {
		newHead = &ListNode{
			Val:  1,
			Next: newHead,
		}
	}

	// reverse the list again
	return newHead
}

func reverse(l *ListNode) *ListNode {
	var prev *ListNode
	n := l
	for n != nil {
		next := n.Next
		n.Next = prev
		prev = n
		n = next
	}

	return prev
}
