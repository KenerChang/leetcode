package populatingnextrightpointersineachnodeii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelMap := map[int]*Node{}
	return connectImpl(root, 0, levelMap)
}

func connectImpl(root *Node, level int, levelMap map[int]*Node) *Node {
	if root == nil {
		return root
	}

	connectImpl(root.Left, level+1, levelMap)
	connectImpl(root.Right, level+1, levelMap)

	if node, found := levelMap[level]; found {
		node.Next = root
	}

	levelMap[level] = root
	return root
}
