package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeImpl(p, q)
}

func isSameTreeImpl(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreeImpl(p.Left, q.Left) && isSameTreeImpl(p.Right, q.Right)
}
