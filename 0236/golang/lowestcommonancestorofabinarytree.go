package lowestcommonancestorofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	paths := map[int]*TreeNode{}

	lowestCommonAncestorRecursive(root, p, root, paths)

	// fmt.Println(paths)

	result, _ := lowestCommonAncestorRecursive(root, q, root, paths)
	// fmt.Println(paths)
	return result
}

func lowestCommonAncestorRecursive(root, target, ancestor *TreeNode, paths map[int]*TreeNode) (*TreeNode, bool) {
	if root == nil {
		return ancestor, false
	}

	if root.Val == target.Val {
		if _, ok := paths[root.Val]; ok {
			ancestor = root
		}
		paths[root.Val] = root

		return ancestor, true
	}

	if _, ok := paths[root.Val]; ok {
		ancestor = root
	}

	paths[root.Val] = root

	result, ok := lowestCommonAncestorRecursive(root.Left, target, ancestor, paths)
	if ok {
		return result, ok
	}

	result, ok = lowestCommonAncestorRecursive(root.Right, target, ancestor, paths)
	if ok {
		return result, ok
	}

	delete(paths, root.Val)
	return ancestor, false
}
