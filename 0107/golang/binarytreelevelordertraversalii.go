package binarytreelevelordertraversalii

// import (
// 	"fmt"
// )

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	result = levelOrderBottomImpl(root, 0, result)

	// revert result
	resultSize := len(result) - 1
	for i := 0; i <= resultSize/2; i++ {
		result[i], result[resultSize-i] = result[resultSize-i], result[i]
	}
	return result
}

func levelOrderBottomImpl(root *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}

	if level >= len(result) {
		result = append(result, []int{})
	}

	result = levelOrderBottomImpl(root.Left, level+1, result)
	result = levelOrderBottomImpl(root.Right, level+1, result)

	result[level] = append(result[level], root.Val)
	return result
}
