package populatingnextrightpointersineachnode

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	posIdx := 1
	levelLastNodeIdx := 1

	queue := []*Node{}
	curNode := root
	var nextNode *Node
	for curNode != nil {
		if curNode.Left != nil && curNode.Right != nil {
			queue = append(queue, curNode.Left, curNode.Right)
		}

		if len(queue) > 0 {
			nextNode = queue[0]
			queue = queue[1:]

			if posIdx != levelLastNodeIdx {
				curNode.Next = nextNode
				fmt.Printf("curNode.Val: %d\n", curNode.Val)
				fmt.Printf("nextNode.Val: %d\n", nextNode.Val)
			} else {
				levelLastNodeIdx = (levelLastNodeIdx << 1) + 1
			}
			curNode = nextNode
			posIdx++
		} else {
			curNode = nil
		}
	}
	return root
}
