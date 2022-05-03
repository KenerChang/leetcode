package binarytreeupsidedown

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	// Do DFS, post order traversal

	stack := []*TreeNode{}
	node := root
	var newRoot *TreeNode
	var newParent *TreeNode

	for node != nil || len(stack) > 0 {
		if node != nil {
			// if node is not nil, put it to stack
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			stack = append(stack, node)

			node = node.Left
		} else {
			// check if right substree is traveled
			node = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			// if is leftest node
			if len(stack) > 0 && node.Right == stack[len(stack)-1] {
				// has not traveled right subtree
				stack = stack[0 : len(stack)-1]
				stack = append(stack, node)

				node = node.Right

			} else {
				// has traveled right subtree or has not right subtree

				// check if the leftmost node
				if newRoot == nil {
					newParent = node
					newRoot = newParent

					node.Left = nil
					node.Right = nil
				} else {
					if node.Left == newParent {
						newParent.Right = node
						newParent = node
					} else {
						newParent.Left = node
					}

					node.Left = nil
					node.Right = nil
				}

				node = nil
			}
		}
	}
	return newRoot
}
