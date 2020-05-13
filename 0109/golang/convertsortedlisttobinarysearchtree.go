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

	nodeMap := map[int]*ListNode{}
	curNode := head
	idx := 0
	for curNode != nil {
		nodeMap[idx] = curNode
		idx++
		curNode = curNode.Next
	}

	return sortedListToBSTImpl(nodeMap, 0, idx)
}

func sortedListToBSTImpl(nodeMap map[int]*ListNode, start, end int) *TreeNode {
	nodeCount := end - start
	if nodeCount <= 0 {
		return nil
	}

	if nodeCount == 1 {
		node := nodeMap[start]
		return &TreeNode{
			Val: node.Val,
		}
	}

	middle := (start + end) / 2
	root := nodeMap[middle]
	return &TreeNode{
		Val:   root.Val,
		Left:  sortedListToBSTImpl(nodeMap, start, middle),
		Right: sortedListToBSTImpl(nodeMap, middle+1, end),
	}
}
