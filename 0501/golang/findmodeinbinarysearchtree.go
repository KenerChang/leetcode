package findmodeinbinarysearchtree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{}
	lastVal := math.MaxInt64
	maxCount := 0
	curCount := 0

	curNode := root

	for len(stack) > 0 || curNode != nil {
		stackSize := len(stack)
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			curNode = stack[stackSize-1]
			stack = stack[:stackSize-1]

			if curNode.Val == lastVal {
				curCount++
			} else {
				// number change
				if curCount > maxCount {
					result = []int{lastVal}
					maxCount = curCount
				} else if curCount == maxCount {
					result = append(result, lastVal)
				}

				lastVal = curNode.Val
				curCount = 1
			}

			curNode = curNode.Right
		}
	}

	if curCount > maxCount {
		result = []int{lastVal}
	} else if curCount == maxCount {
		result = append(result, lastVal)
	}

	return result
}
