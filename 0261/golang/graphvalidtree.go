package graphvalidtree

func validTree(n int, edges [][]int) bool {
	// if len(edges) is not n-1, it is not a tree
	if len(edges) != n-1 {
		return false
	}

	// tree condition
	//  - no cycle
	//  - no isolated node

	// use a slice to keep parent
	parents := initParents(n)

	// start from any node
	for _, edge := range edges {
		// get parent of startNode, and endNode
		// the two node should not have same root, or the graph has cycle

		startNode := edge[0]
		endNode := edge[1]

		parentOfStart := getParent(parents, startNode)
		parentOfEnd := getParent(parents, endNode)

		if parentOfStart == parentOfEnd {
			return false
		}

		union(parents, parentOfStart, startNode, endNode)
	}

	// if there are nodes which are not visited, it is not a tree
	return true
}

func initParents(n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	return result
}

func getParent(parents []int, node int) int {
	if parents[node] == -1 {
		return node
	}

	return getParent(parents, parents[node])
}

func union(parents []int, parentStart, start, end int) {
	var node int
	if parentStart == start {
		node = start
	} else {
		node = parentStart
	}
	parents[node] = end
}
