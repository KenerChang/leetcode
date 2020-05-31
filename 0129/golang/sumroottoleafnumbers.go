package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersImpl(root, 0)
}

func sumNumbersImpl(root *TreeNode, accumulated int) int {
	if root == nil {
		return 0
	}

	nowAccumulated := accumulated*10 + root.Val
	leftChild := sumNumbersImpl(root.Left, nowAccumulated)
	rightChild := sumNumbersImpl(root.Right, nowAccumulated)

	// leaf node
	if leftChild == 0 && rightChild == 0 {
		return nowAccumulated
	}

	return leftChild + rightChild
}
