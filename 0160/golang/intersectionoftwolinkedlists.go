package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var countA, countB int
	var nodeA, nodeB *ListNode = headA, headB
	for nodeA != nil || nodeB != nil {
		if nodeA != nil {
			countA++
			nodeA = nodeA.Next
		}

		if nodeB != nil {
			countB++
			nodeB = nodeB.Next
		}
	}

	nodeA, nodeB = headA, headB
	for nodeA != nodeB {
		if countA > countB {
			nodeA = nodeA.Next
			countA--
			continue
		}

		if countB > countA {
			nodeB = nodeB.Next
			countB--
			continue
		}

		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nodeA
}
