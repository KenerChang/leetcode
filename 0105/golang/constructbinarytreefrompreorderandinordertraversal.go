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
	leftPreorder, rightPreorder := splitPreorderTree(preorder, leftInorder, rightInorder)

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

func splitPreorderTree(preorder, leftInorder, rightInorder []int) ([]int, []int) {
	leftMap := map[int]bool{}
	for _, num := range leftInorder {
		leftMap[num] = true
	}

	rightMap := map[int]bool{}
	for _, num := range rightInorder {
		rightMap[num] = true
	}

	leftPreorder := []int{}
	rightPreorder := []int{}
	for _, num := range preorder {
		if _, found := leftMap[num]; found {
			leftPreorder = append(leftPreorder, num)
			continue
		}

		if _, found := rightMap[num]; found {
			rightPreorder = append(rightPreorder, num)
			continue
		}
	}
	return leftPreorder, rightPreorder
}
