package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Vals() []int {
	results := []int{}
	head := l

	for head != nil {
		results = append(results, head.Val)
		head = head.Next
	}

	return results
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	count := 0
	var newHead, lastTail *ListNode
	start, end, ptr := head, head, head
	for ptr != nil {
		if count%k == k-1 && count > 0 {
			end = ptr
			ptr = ptr.Next

			reversedHead, reversedTail := doReverse(start, end)

			reversedTail.Next = ptr
			if lastTail != nil {
				lastTail.Next = reversedHead
			} else {
				newHead = reversedHead
			}
			lastTail = reversedTail
		} else if count%k == 0 {
			start = ptr
			ptr = ptr.Next
		} else {
			ptr = ptr.Next
		}
		count++
	}

	return newHead
}

func doReverse(head, tail *ListNode) (newHead, newTail *ListNode) {
	prev, next := head, head.Next

	newHead, newTail = head, head
	for newHead != tail {
		prev = newHead
		newHead = next
		next = next.Next
		newHead.Next = prev
	}
	return
}
