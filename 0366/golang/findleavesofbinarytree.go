package findleavesofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func removeChildAndReturnItsDistance(root *TreeNode, results [][]int) (int, [][]int) {
	if root.Left == nil && root.Right == nil {
		if len(results) == 0 {
			results = append(results, []int{})
		}

		results[0] = append(results[0], root.Val)
		return 0, results
	}

	var leftDistance int
	if root.Left != nil {
		leftDistance, results = removeChildAndReturnItsDistance(root.Left, results)
	}

	var rightDistance int
	if root.Right != nil {
		rightDistance, results = removeChildAndReturnItsDistance(root.Right, results)
	}

	distance := max(leftDistance, rightDistance) + 1
	if len(results) <= distance {
		results = append(results, []int{root.Val})
	} else {
		results[distance] = append(results[distance], root.Val)
	}

	// remove children
	root.Left = nil
	root.Right = nil

	return distance, results
}

func findLeaves(root *TreeNode) [][]int {
	// we mark each node with distance from leaf
	// then we know where the node should appear in result

	// To know the node distance, we need a util function recursively
	// to find th child distance and remove the child
	_, results := removeChildAndReturnItsDistance(root, [][]int{})
	return results
}
