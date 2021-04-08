package kthsmallestelementinabst

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func kthSmallest(root *TreeNode, k int) int {
	_, result := kthSmallestRecursive(root, k, -1)
	return result
}

func kthSmallestRecursive(root *TreeNode, remain, result int) (int, int) {
	if remain == 0 || root == nil {
		return remain, result
	}

	remain, result = kthSmallestRecursive(root.Left, remain, result)
	if remain <= 0 {
		return remain, result
	}

	if remain == 1 {
		return remain - 1, root.Val
	}
	return kthSmallestRecursive(root.Right, remain-1, result)
}
