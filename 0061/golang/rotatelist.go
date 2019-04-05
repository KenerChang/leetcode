package rotatelist

// import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// get nodes number in the list
	total := 0
	cNode := head
	var oldLastNode *ListNode
	for cNode != nil {
		total++

		// last node
		if cNode.Next == nil {
			oldLastNode = cNode
			break
		}
		cNode = cNode.Next
	}

	targetRemainder := k % total
	if targetRemainder == 0 {
		return head
	}

	count := 0
	var newHead *ListNode
	var newLastNode *ListNode
	cNode = head
	for cNode != nil {
		reverseRemainder := total - count
		if reverseRemainder == targetRemainder+1 {
			newLastNode = cNode

		} else if reverseRemainder == targetRemainder {
			newHead = cNode
			break
		}
		count++
		cNode = cNode.Next
	}
	oldLastNode.Next = head
	newLastNode.Next = nil
	return newHead
}
