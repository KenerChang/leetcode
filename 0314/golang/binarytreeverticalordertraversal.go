package binarytreeverticalordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type wrapTreeNode struct {
	Node *TreeNode
	X    int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func verticalOrder(root *TreeNode) [][]int {
	// We can use BFS to know the left & right boundary
	if root == nil {
		return [][]int{}
	}

	var l, r int
	queue := []*wrapTreeNode{
		&wrapTreeNode{
			Node: root,
			X:    0,
		},
	}

	// do bfs
	bfsResults := [][]int{} // 0: val, 1: pos
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		bfsResults = append(bfsResults, []int{node.Node.Val, node.X})
		l = min(l, node.X)
		r = max(r, node.X)

		if node.Node.Left != nil {
			queue = append(queue, &wrapTreeNode{
				Node: node.Node.Left,
				X:    node.X - 1,
			})
		}

		if node.Node.Right != nil {
			queue = append(queue, &wrapTreeNode{
				Node: node.Node.Right,
				X:    node.X + 1,
			})
		}
	}

	// After we get left & right boundary
	// we can iterate the tree again and get where the node should belong to
	results := make([][]int, r-l+1)

	for _, result := range bfsResults {
		// need to shift x-axis with l
		pos := result[1] - l

		if results[pos] == nil {
			results[pos] = []int{}
		}

		results[pos] = append(results[pos], result[0])
	}

	return results
}
