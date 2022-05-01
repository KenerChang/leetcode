package binarytreemaximumpathsum

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// use DFS to accumulate subtree sum.
	// we use post order to travel the tree
	// so each tree sum can be node.Val + node.Left.Val + node.Right.Val

	return maxPathSumRecursive(root, math.MinInt)
}

func maxPathSumRecursive(root *TreeNode, currMax int) int {
	if root == nil {
		return currMax
	}

	// if root is leaf, return root.Val
	if root.Left == nil && root.Right == nil {
		return max(root.Val, currMax)
	}
	// fmt.Printf("root.Val before: %d\n", root.Val)

	// go left
	currMax = maxPathSumRecursive(root.Left, currMax)

	// go right
	currMax = maxPathSumRecursive(root.Right, currMax)

	// check new path
	newPath := root.Val
	if root.Left != nil {
		newPath = max(newPath, newPath+root.Left.Val)
	}

	if root.Right != nil {
		newPath = max(newPath, newPath+root.Right.Val)
	}

	currMax = max(currMax, newPath)

	// set node vaule to max gain
	root.Val = max(root.Val, root.Val+maxGain(root))
	// fmt.Printf("maxGain: %d\n", maxGain(root))
	// fmt.Printf("root.Val after: %d\n", root.Val)

	return currMax
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxGain(root *TreeNode) int {
	leftGain := math.MinInt
	rightGain := math.MinInt

	if root.Left != nil {
		leftGain = root.Left.Val
	}

	if root.Right != nil {
		rightGain = root.Right.Val
	}

	return max(leftGain, rightGain)
}
