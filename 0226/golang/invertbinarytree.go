package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := root.Right

	root.Right = root.Left
	root.Left = temp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
