package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	visited := map[*ListNode]bool{}
	currNode := head

	for currNode != nil {
		if _, found := visited[currNode]; found {
			return true
		}
		visited[currNode] = true
		currNode = currNode.Next
	}
	return false
}
