package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cNode := head // current node
	var newHead *ListNode
	var cLast *ListNode // current last
	var totalLast *ListNode
	var lfirst *ListNode

	for cNode != nil {
		cNextNode := cNode.Next
		if cNode.Val < x {
			if newHead == nil {
				newHead = cNode
				cLast = newHead
				if cNode != head {
					cNode.Next = head

					if totalLast != nil {
						totalLast.Next = cNextNode
					}
				}
			} else {
				next := cLast.Next
				cLast.Next = cNode
				cNode.Next = next
				cLast = cNode

				if totalLast != nil {
					totalLast.Next = cNextNode
				}
			}

			if lfirst != nil {
				cLast.Next = lfirst
			} else {
				cLast.Next = cNextNode
			}
		} else {
			if lfirst == nil {
				lfirst = cNode
			}
			totalLast = cNode
		}
		cNode = cNextNode
	}
	if newHead == nil {
		newHead = head
	}

	if totalLast != nil {
		totalLast.Next = nil
	} else {
		cLast.Next = nil
	}
	return newHead
}
