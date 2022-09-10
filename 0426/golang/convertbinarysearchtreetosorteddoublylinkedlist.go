package convertbinarysearchtreetosorteddoublylinkedlist

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) ValueList() []int {
	numMap := map[int]bool{}

	root := n
	nums := []int{}
	for !numMap[root.Val] {
		numMap[root.Val] = true
		nums = append(nums, root.Val)
		root = root.Right
	}

	return nums
}

func treeToDoublyList(root *Node) *Node {
	// preoder traversal

	if root == nil {
		return nil
	}

	var newHead, prev *Node

	stack := []*Node{}
	n := root

	for n != nil || len(stack) > 0 {
		if n != nil {
			stack = append(stack, n)
			n = n.Left
		} else {
			// pop from stack

			n = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if newHead == nil {
				// set new head to this node
				newHead = n
			} else {
				// link prev to this node
				prev.Right = n
				n.Left = prev
			}
			prev = n

			n = n.Right
		}
	}

	prev.Right = newHead
	newHead.Left = prev

	return newHead
}
