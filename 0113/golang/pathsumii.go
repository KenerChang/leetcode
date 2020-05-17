package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	return pathSumImpl(root, sum, [][]int{}, []int{})
}

func pathSumImpl(root *TreeNode, sum int, results [][]int, path []int) [][]int {
	if root == nil {
		return results
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		path = append(path, root.Val)

		results = append(results, path)
		return results
	}

	remain := sum - root.Val
	copiedPath := append([]int{}, path...)
	copiedPath = append(copiedPath, root.Val)
	results = pathSumImpl(root.Left, remain, results, copiedPath)

	copiedPath = append([]int{}, path...)
	copiedPath = append(copiedPath, root.Val)
	results = pathSumImpl(root.Right, remain, results, copiedPath)
	return results
}
