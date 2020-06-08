package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slowP := head
	fastP := head
	hasCycle := false
	for fastP.Next != nil && fastP.Next.Next != nil {
		fastP = fastP.Next.Next
		slowP = slowP.Next
		if slowP == fastP {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	for slowP != head {
		slowP = slowP.Next
		head = head.Next
	}
	return slowP
}
