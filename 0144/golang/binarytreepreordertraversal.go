package binarytreepreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{}
	node := root
	results := []int{}
	for node != nil || len(stack) > 0 {
		if node == nil {
			last := len(stack) - 1
			node = stack[last]
			stack = stack[0:last]
		}
		results = append(results, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		node = node.Left
	}
	return results
}
