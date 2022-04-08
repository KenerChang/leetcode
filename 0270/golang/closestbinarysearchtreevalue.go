package closestbinarysearchtreevalue

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	node := root
	minVal := math.MaxFloat64
	var nodeVal int
	// start from root
	for node != nil {
		diff := float64(node.Val) - target
		diff = math.Abs(diff)

		if diff < minVal {
			// if node.Val - target < minVal, update minVal & nodeVal
			minVal = diff
			nodeVal = node.Val
		} else if minVal == 0 {
			// else if node.Val - target > minVal, no need to go down since the difference will become larger
			break
		}

		// if target is larger than root go right, else go left
		if float64(node.Val) < target {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return nodeVal
}
