package verticalordertraversalofabinarytree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type wrappedTreeNode struct {
	Node *TreeNode
	X    int
	Y    int
}

type byXAndByVal [][]int

func (b byXAndByVal) Len() int {
	return len(b)
}

func (b byXAndByVal) Less(i, j int) bool {
	if b[i][1] != b[j][1] {
		return b[i][1] < b[j][1]
	}

	if b[i][2] != b[j][2] {
		return b[i][2] < b[j][2]
	}

	return b[i][0] < b[j][0]
}

func (b byXAndByVal) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
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

func verticalTraversal(root *TreeNode) [][]int {
	// Use BFS to traverse the tree

	if root == nil {
		return [][]int{}
	}

	queue := []*wrappedTreeNode{
		{
			Node: root,
			X:    0,
		},
	}

	bfsResults := [][]int{}
	var l, r int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		bfsResults = append(bfsResults, []int{node.Node.Val, node.X, node.Y})
		l = min(l, node.X)
		r = max(r, node.X)

		if node.Node.Left != nil {
			queue = append(queue, &wrappedTreeNode{
				Node: node.Node.Left,
				X:    node.X - 1,
				Y:    node.Y + 1,
			})
		}

		if node.Node.Right != nil {
			queue = append(queue, &wrappedTreeNode{
				Node: node.Node.Right,
				X:    node.X + 1,
				Y:    node.Y + 1,
			})
		}
	}

	sort.Sort(byXAndByVal(bfsResults))

	// output results
	results := make([][]int, r-l+1)

	for _, result := range bfsResults {
		x := result[1] - l
		if results[x] == nil {
			results[x] = []int{}
		}

		results[x] = append(results[x], result[0])
	}

	return results
}
