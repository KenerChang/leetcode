package symmetrictree

import (
// "fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	leftNode := root.Left
	rightNode := root.Right
	leftTreeStack := []*TreeNode{}
	rightTreeStack := []*TreeNode{}

	var leftStackSize, rightStackSize int

	for (leftNode != nil || len(leftTreeStack) > 0) && (rightNode != nil || len(rightTreeStack) > 0) {
		if (leftNode != nil && rightNode == nil) || (leftNode == nil && rightNode != nil) {
			// one is nil, another is not
			return false
		}

		if leftNode != nil && rightNode != nil {
			// both are not nil
			leftTreeStack = append(leftTreeStack, leftNode)
			rightTreeStack = append(rightTreeStack, rightNode)

			leftNode = leftNode.Left
			rightNode = rightNode.Right

		} else {
			// both are nil
			leftStackSize = len(leftTreeStack)
			leftNode = leftTreeStack[leftStackSize-1]
			leftTreeStack = leftTreeStack[:leftStackSize-1]

			rightStackSize = len(rightTreeStack)
			rightNode = rightTreeStack[rightStackSize-1]
			rightTreeStack = rightTreeStack[:rightStackSize-1]

			if leftNode.Val != rightNode.Val {
				return false
			}

			leftNode = leftNode.Right
			rightNode = rightNode.Left
		}
	}

	if len(leftTreeStack) != len(rightTreeStack) {
		return false
	}

	if (leftNode != nil && rightNode == nil) || (leftNode == nil && rightNode != nil) {
		// one is nil, another is not
		return false
	}

	if leftNode != nil && rightNode != nil {
		return leftNode.Val == rightNode.Val
	}
	return true
}
