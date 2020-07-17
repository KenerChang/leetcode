package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	levels := [][]int{}
	levels = rightSideViewImpl(root, 0, levels)

	result := []int{}
	for _, level := range levels {
		result = append(result, level[0])
	}
	return result
}

func rightSideViewImpl(root *TreeNode, level int, levels [][]int) [][]int {
	if root == nil {
		return levels
	}

	if level >= len(levels) {
		levels = append(levels, []int{root.Val})
	}

	levels = rightSideViewImpl(root.Right, level+1, levels)
	levels = rightSideViewImpl(root.Left, level+1, levels)

	return levels
}
