package binarytreepruning

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsUtil(root *TreeNode) bool {
	// an util function which iterates trees, remove subtree if not containing 1
	// and return if it has at least one 1 in its subtrees

	if root == nil {
		return false
	}

	leftHasOne := dfsUtil(root.Left)
	if !leftHasOne {
		root.Left = nil
	}

	rightHasOne := dfsUtil(root.Right)
	if !rightHasOne {
		root.Right = nil
	}

	return root.Val == 1 || leftHasOne || rightHasOne

}

func pruneTree(root *TreeNode) *TreeNode {
	// Can use dfs to iterate subtrees and do the work
	hasOne := dfsUtil(root)
	if !hasOne {
		return nil
	}

	return root
}
