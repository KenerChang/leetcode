package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowP := head
	fastP := head.Next
	for slowP != fastP {
		if fastP == nil || fastP.Next == nil {
			return false
		}
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
	return true
}
