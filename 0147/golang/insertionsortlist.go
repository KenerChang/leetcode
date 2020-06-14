package insertionsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	node := head

	for node != nil {
		next := node.Next
		if node == head {
			node = next
			head.Next = nil
			continue
		}

		var prev *ListNode
		ptr := newHead
		for ptr != nil {
			if node.Val > ptr.Val {
				prev = ptr
				ptr = ptr.Next
				continue
			}

			if ptr == newHead {
				newHead = node
			} else {
				prev.Next = node
			}

			node.Next = ptr
			break
		}

		if ptr == nil {
			prev.Next = node
			node.Next = nil
		}

		node = next
	}
	return newHead
}
