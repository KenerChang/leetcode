package binarytreepostordertraversal

// import (
// 	"fmt"
// )

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	results := []int{}
	if root == nil {
		return results
	}

	stack := []*TreeNode{}
	node := root
	var prev *TreeNode

	for node != nil || len(stack) > 0 {
		if node == nil {
			size := len(stack) - 1
			node = stack[size]
			stack = stack[0:size]
		}

		if (prev == node.Right && node.Right != nil) || (prev == node.Left && node.Right == nil) || (node.Left == nil && node.Right == nil) {
			results = append(results, node.Val)
			prev = node
			node = nil
			continue
		}

		stack = append(stack, node)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		prev = node
		node = node.Left
	}

	return results
}
