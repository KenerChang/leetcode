package countunivaluesubtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	result, _ := countUnivalSubtreesRecursive(root, 0)
	return result
}

func countUnivalSubtreesRecursive(root *TreeNode, accumulated int) (int, bool) {
	// if root is nil, return 0
	if root == nil {
		return accumulated, true
	}

	// if root is leaf, increase accumulated by 1
	if root.Left == nil && root.Right == nil {
		return accumulated + 1, true
	}

	// if left child & right child are unival subtrees, and root.Val is equal to left.val & right.val
	// root is also a unival subtree, increase accumulated by 1
	var isLeftUnival, isRightUnival bool
	accumulated, isLeftUnival = countUnivalSubtreesRecursive(root.Left, accumulated)
	accumulated, isRightUnival = countUnivalSubtreesRecursive(root.Right, accumulated)

	if isLeftUnival && isRightUnival {
		if root.Left != nil && root.Val != root.Left.Val {
			return accumulated, false
		}

		if root.Right != nil && root.Val != root.Right.Val {
			return accumulated, false
		}

		return accumulated + 1, true
	}

	return accumulated, false

}
