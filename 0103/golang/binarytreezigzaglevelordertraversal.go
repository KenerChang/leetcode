package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	return zigzagLevelOrderImpl(root, 0, result)
}

func zigzagLevelOrderImpl(root *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}

	if level >= len(result) {
		result = append(result, []int{})
	}

	if level%2 == 1 {
		result[level] = append([]int{root.Val}, result[level]...)
	} else {
		result[level] = append(result[level], root.Val)
	}

	result = zigzagLevelOrderImpl(root.Left, level+1, result)
	result = zigzagLevelOrderImpl(root.Right, level+1, result)
	return result
}
