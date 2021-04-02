package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodesRecursive(root, 1)
}

func countNodesRecursive(node *TreeNode, number int) int {
	if node.Left == nil && node.Right == nil {
		return number
	}

	if node.Right == nil {
		return 2 * number
	}

	return max(countNodesRecursive(node.Left, 2*number), countNodesRecursive(node.Right, 2*number+1))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
