package removeduplicatedfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// we iterate the list and check if a number is duplicated
	// if a number is not duplicated, it will add to the last of new list
	// and it will be the new last node of new list
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var newHead *ListNode
	var currLast *ListNode

	duplicate := false
	for curr != nil {
		next := curr.Next
		if curr.Next == nil || curr.Val != next.Val {
			if !duplicate {
				if newHead == nil {
					newHead = curr
				}

				if currLast == nil {
					currLast = curr
				} else {
					currLast.Next = curr
					currLast = curr
				}
			}
			duplicate = false
		} else {
			duplicate = true
		}
		curr = curr.Next
	}
	if currLast != nil {
		currLast.Next = nil
	}
	return newHead
}
