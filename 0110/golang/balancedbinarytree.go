package balancedbinarytree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, err := isBalancedImpl(root)
	return err == nil
}

func isBalancedImpl(root *TreeNode) (int, error) {
	if root == nil {
		return 0, nil
	}

	if root.Left == nil && root.Right == nil {
		return 1, nil
	}

	leftTreeHieght, err := isBalancedImpl(root.Left)
	if err != nil {
		return 0, err
	}

	rightTreeHeight, err := isBalancedImpl(root.Right)
	if err != nil {
		return 0, err
	}

	diff := leftTreeHieght - rightTreeHeight
	if diff >= 2 || diff <= -2 {
		return 0, fmt.Errorf("error")
	}

	return 1 + max(leftTreeHieght, rightTreeHeight), nil
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}
