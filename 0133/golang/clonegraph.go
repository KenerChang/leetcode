package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[int]*Node{}
	return cloneGraphImpl(node, visited)
}

func cloneGraphImpl(node *Node, visitedCache map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if copied, found := visitedCache[node.Val]; found {
		return copied
	}

	copied := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visitedCache[node.Val] = copied

	for i, neighbor := range node.Neighbors {
		copiedNeighbor := cloneGraphImpl(neighbor, visitedCache)
		copied.Neighbors[i] = copiedNeighbor
	}
	return copied
}
