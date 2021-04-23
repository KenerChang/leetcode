package deletenodeinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToSlice() []int {
	results := []int{}

	curr := l
	for curr != nil {
		results = append(results, curr.Val)
		curr = curr.Next
	}
	return results
}

func deleteNode(node *ListNode) {
	curr := node.Next
	prev := node

	for curr != nil {

		prev.Val = curr.Val
		if curr.Next != nil {
			prev = curr
			curr = curr.Next
		} else {

			break
		}
	}

	prev.Next = nil
}
