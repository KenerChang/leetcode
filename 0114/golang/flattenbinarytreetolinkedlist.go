package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	stack := []*TreeNode{}

	curNode := root
	for curNode != nil {
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}

		if curNode.Left != nil {
			curNode.Right = curNode.Left
			curNode.Left = nil
		} else {
			stackSize := len(stack) - 1
			if stackSize >= 0 {
				nextNode := stack[stackSize]
				curNode.Right = nextNode
				stack = stack[0:stackSize]
			} else {
				curNode.Right = nil
			}
		}

		curNode = curNode.Right
	}
}
