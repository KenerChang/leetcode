package constructbinarytreefrominorderandpostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for idx, num := range inorder {
		inorderMap[num] = idx
	}

	return buildTreeImpl(inorder, postorder, inorderMap, 0)
}

func buildTreeImpl(inorder, postorder []int, inorderMap map[int]int, offset int) *TreeNode {
	treeSize := len(postorder)
	if treeSize == 0 {
		return nil
	}

	root := postorder[treeSize-1]

	// LVR
	inorderRootIdx := inorderMap[root] - offset
	inorderLeftTree := inorder[0:inorderRootIdx]
	inorderRightTree := inorder[:inorderRootIdx+1]

	// LRV
	postorderLeftTree := postorder[0:len(inorderLeftTree)]
	postorderRightTree := postorder[len(inorderLeftTree) : treeSize-1]

	tree := &TreeNode{
		Val:   root,
		Left:  buildTreeImpl(inorderLeftTree, postorderLeftTree, inorderMap, offset),
		Right: buildTreeImpl(inorderRightTree, postorderRightTree, inorderMap, offset+inorderRootIdx+1),
	}

	return tree
}
