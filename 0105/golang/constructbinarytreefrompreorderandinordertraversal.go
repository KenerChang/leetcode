package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for idx, num := range inorder {
		inorderMap[num] = idx
	}

	return buildTreeImpl(preorder, inorder, inorderMap, 0)
}

func buildTreeImpl(preorder []int, inorder []int, inorderMap map[int]int, offset int) *TreeNode {
	// preoder: VLR
	// inorder: LVR

	if (len(preorder) == 0 && len(inorder) == 0) || (len(preorder) != len(inorder)) {
		return nil
	}

	root := preorder[0]
	idx := findIdx(inorderMap, root, offset)
	leftInorder, rightInorder := splitInorderTree(inorder, idx)
	leftPreorder, rightPreorder := splitPreorderTree(preorder, len(leftInorder), len(rightInorder))

	tree := &TreeNode{
		Val:   root,
		Left:  buildTreeImpl(leftPreorder, leftInorder, inorderMap, offset),
		Right: buildTreeImpl(rightPreorder, rightInorder, inorderMap, offset+idx+1),
	}

	return tree
}

func findIdx(idxMap map[int]int, target, offset int) int {
	idx, _ := idxMap[target]
	return idx - offset
}

func splitInorderTree(inorder []int, rootIdx int) ([]int, []int) {
	treeSize := len(inorder)
	if rootIdx == 0 {
		return []int{}, inorder[1:treeSize]
	} else if rootIdx == treeSize {
		return inorder[0 : treeSize-1], []int{}
	}

	return inorder[0:rootIdx], inorder[rootIdx+1 : treeSize]
}

func splitPreorderTree(preorder []int, leftTreeSize, rightTreeSize int) ([]int, []int) {
	treeSize := len(preorder)
	if leftTreeSize == 0 {
		return []int{}, preorder[1:treeSize]
	} else if rightTreeSize == 0 {
		return preorder[1:treeSize], []int{}
	}
	return preorder[1 : leftTreeSize+1], preorder[leftTreeSize+1 : treeSize]
}
