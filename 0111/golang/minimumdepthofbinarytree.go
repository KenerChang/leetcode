package minimumdepthofbinarytree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return minDepthImpl(root)
}

func minDepthImpl(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := math.MaxInt64
	if root.Left != nil {
		leftDepth = minDepthImpl(root.Left)
	}

	rightDepth := math.MaxInt64
	if root.Right != nil {
		rightDepth = minDepthImpl(root.Right)
	}

	return 1 + min(leftDepth, rightDepth)
}

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
