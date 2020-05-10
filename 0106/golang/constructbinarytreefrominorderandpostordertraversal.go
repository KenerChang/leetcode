package constructbinarytreefrominorderandpostordertraversal

// import (
// 	"fmt"
// )

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

	return buildTreeImpl(postorder, inorderMap, 0)
}

func buildTreeImpl(postorder []int, inorderMap map[int]int, offset int) *TreeNode {
	treeSize := len(postorder)
	if treeSize == 0 {
		return nil
	}

	root := postorder[treeSize-1]

	// LVR
	inorderRootIdx := inorderMap[root] - offset

	// LRV
	postorderLeftTree := postorder[0:inorderRootIdx]
	postorderRightTree := postorder[len(postorderLeftTree) : treeSize-1]

	tree := &TreeNode{
		Val:   root,
		Left:  buildTreeImpl(postorderLeftTree, inorderMap, offset),
		Right: buildTreeImpl(postorderRightTree, inorderMap, offset+inorderRootIdx+1),
	}

	return tree
}
