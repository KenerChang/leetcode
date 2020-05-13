package convertsortedlisttobinarysearchtree

// import (
// 	"fmt"
// )

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	curNode := head
	nodeCount := 0
	for curNode != nil {
		nodeCount++
		curNode = curNode.Next
	}

	return sortedListToBSTImpl(head, nodeCount)
}

func sortedListToBSTImpl(start *ListNode, nodeCount int) *TreeNode {
	if start == nil || nodeCount == 0 {
		return nil
	}

	if nodeCount == 1 {
		return &TreeNode{
			Val: start.Val,
		}
	}

	idx := 0
	middle := nodeCount / 2
	curNode := start

	var tree *TreeNode
	for idx <= middle+1 && curNode != nil {
		if idx == middle {
			tree = &TreeNode{
				Val:  curNode.Val,
				Left: sortedListToBSTImpl(start, middle),
			}
		}

		if idx == middle+1 {
			tree.Right = sortedListToBSTImpl(curNode, nodeCount-middle-1)
		}
		idx++
		curNode = curNode.Next
	}
	return tree
}
