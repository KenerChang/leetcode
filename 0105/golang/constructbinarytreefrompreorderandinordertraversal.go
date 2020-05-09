package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// preoder: VLR
	// inorder: LVR

	if (len(preorder) == 0 && len(inorder) == 0) || (len(preorder) != len(inorder)) {
		return nil
	}

	root := preorder[0]
	idx := findIdx(inorder, root)
	leftInorder, rightInorder := splitInorderTree(inorder, idx)
	leftPreorder, rightPreorder := splitPreorderTree(preorder, len(leftInorder), len(rightInorder))

	tree := &TreeNode{
		Val:   root,
		Left:  buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}

	return tree
}

func findIdx(nums []int, target int) int {
	for idx, num := range nums {
		if num == target {
			return idx
		}
	}
	return -1
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
