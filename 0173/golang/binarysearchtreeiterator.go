package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		Stack: []*TreeNode{},
	}

	iterator.push(root)
	return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.pop()
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

func (this *BSTIterator) pop() *TreeNode {
	if len(this.Stack) == 0 {
		return nil
	}

	last := len(this.Stack) - 1
	node := this.Stack[last]
	this.Stack = this.Stack[0:last]

	if node.Right != nil {
		this.push(node.Right)
	}

	return node
}

func (this *BSTIterator) push(node *TreeNode) {
	if node == nil {
		return
	}

	curNode := node
	for curNode != nil {
		this.Stack = append(this.Stack, curNode)
		curNode = curNode.Left
	}
}
