package validatebinarysearchtree

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	return isValidBSTImpl(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTImpl(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val >= upper || root.Val <= lower {
		return false
	}

	return isValidBSTImpl(root.Left, lower, root.Val) && isValidBSTImpl(root.Right, root.Val, upper)
}
