package binarytreelevelordertraversal

import (
// "fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	results := [][]int{}

	return levelOrderImpl(root, 0, results)
}

func levelOrderImpl(root *TreeNode, level int, levelorder [][]int) [][]int {
	if root == nil {
		return levelorder
	}

	if level >= len(levelorder) {
		levelorder = append(levelorder, []int{})
	}

	levelorder[level] = append(levelorder[level], root.Val)

	levelorder = levelOrderImpl(root.Left, level+1, levelorder)
	levelorder = levelOrderImpl(root.Right, level+1, levelorder)

	return levelorder
}
