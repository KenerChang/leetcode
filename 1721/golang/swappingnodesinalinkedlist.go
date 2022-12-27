package swappingnodesinalinkedlist

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

func swapNodes(head *ListNode, k int) *ListNode {
	// find the k-th node, k-th node's previous node & list size n
	count := 0
	n := head

	var prevNodeOfKth *ListNode
	var kthNode *ListNode
	for n != nil {
		count++

		if count == k-1 {
			prevNodeOfKth = n
		} else if count == k {
			kthNode = n
		}

		n = n.Next
	}
	listSize := count

	// fmt.Printf("count: %d, prevNodeOfKth: %+v, kthNode: %+v\n", count, prevNodeOfKth, kthNode)

	// find the (n-k+1)-th node and its previous node
	n = head
	var prevNodeOfReverseKth *ListNode
	var reverseKthNode *ListNode

	count = 0
	for n != nil {
		count++

		if count == listSize-k {
			prevNodeOfReverseKth = n
		} else if count == listSize-k+1 {
			reverseKthNode = n
		}

		n = n.Next
	}

	// fmt.Printf("count: %d, prevNodeOfReverseKth: %+v, reverseKthNode: %+v\n", count, prevNodeOfReverseKth, reverseKthNode)

	newHead := head

	// do swap
	if prevNodeOfKth != nil {
		prevNodeOfKth.Next = reverseKthNode
	}

	if prevNodeOfReverseKth != nil {
		prevNodeOfReverseKth.Next = kthNode
	}

	kthNodeNext := kthNode.Next
	kthNode.Next = reverseKthNode.Next
	reverseKthNode.Next = kthNodeNext

	if prevNodeOfKth == nil || prevNodeOfReverseKth == nil {
		if k == 1 {
			newHead = reverseKthNode
		} else {
			newHead = kthNode
		}
	}

	return newHead
}
