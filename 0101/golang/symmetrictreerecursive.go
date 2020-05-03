package symmetrictree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	return isSymmetricRecursive(root.Left, root.Right)
}

func isSymmetricRecursive(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		// both nil
		return true
	}

	if tree1 == nil || tree2 == nil {
		// one is nil
		return false
	}

	return (tree1.Val == tree2.Val) && isSymmetricRecursive(tree1.Left, tree2.Right) && isSymmetricRecursive(tree1.Right, tree2.Left)
}
