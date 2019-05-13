package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cNode := head
	nList := head
	for cNode.Next != nil {
		next := cNode.Next

		if cNode.Val != next.Val {
			nList.Next = next
			nList = next
		}
		cNode = next
	}
	nList.Next = nil
	return head
}
