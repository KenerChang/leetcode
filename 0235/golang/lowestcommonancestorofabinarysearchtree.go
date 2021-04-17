package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	paths := map[int]*TreeNode{}

	// find paths of p
	lowestCommonAncestorRecursive(root, p, root, paths)

	lowestAncestor := lowestCommonAncestorRecursive(root, q, root, paths)
	return lowestAncestor
}

func lowestCommonAncestorRecursive(root *TreeNode, target *TreeNode, ancestor *TreeNode, paths map[int]*TreeNode) *TreeNode {
	if root == nil || root.Val == target.Val {
		if _, ok := paths[root.Val]; ok {
			ancestor = root
		} else {
			paths[root.Val] = root
		}

		return ancestor
	}

	if _, ok := paths[root.Val]; ok {
		ancestor = root
	} else {
		paths[root.Val] = root
	}

	var lowestAncestor *TreeNode
	if target.Val > root.Val {
		lowestAncestor = lowestCommonAncestorRecursive(root.Right, target, ancestor, paths)
	} else {
		lowestAncestor = lowestCommonAncestorRecursive(root.Left, target, ancestor, paths)
	}

	return lowestAncestor
}
