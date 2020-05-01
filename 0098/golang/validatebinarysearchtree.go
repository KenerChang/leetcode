package validatebinarysearchtree

import (
	"math/big"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ZERO *big.Int = big.NewInt(0)
var ONE *big.Int = big.NewInt(1)
var TWO *big.Int = big.NewInt(2)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodeNum := big.NewInt(1)
	treePos := map[string]*TreeNode{
		"1": root,
	}
	return isValidBSTImpl(root, treePos, nodeNum)
}

func isValidBSTImpl(tree *TreeNode, treePos map[string]*TreeNode, nodeNum *big.Int) bool {
	if tree == nil {
		return true
	}

	passed := validate(tree, treePos, nodeNum)
	if !passed {
		return false
	}

	treePos[nodeNum.String()] = tree

	leftNodeNum := big.NewInt(1)
	leftNodeNum = leftNodeNum.Mul(nodeNum, TWO)
	passed = isValidBSTImpl(tree.Left, treePos, leftNodeNum)
	if !passed {
		return false
	}

	rightNodeNum := big.NewInt(1)
	rightNodeNum = rightNodeNum.Mul(nodeNum, TWO)
	rightNodeNum = rightNodeNum.Add(rightNodeNum, ONE)

	passed = isValidBSTImpl(tree.Right, treePos, rightNodeNum)
	if !passed {
		return false
	}

	return validate(tree, treePos, nodeNum)
}

func validate(node *TreeNode, treePos map[string]*TreeNode, nodeNum *big.Int) bool {
	ancestorNum := big.NewInt(1)
	remainder := big.NewInt(1)
	ancestorNum, remainder = ancestorNum.DivMod(nodeNum, TWO, remainder)

	for ancestorNum.Cmp(ZERO) == 1 {
		ancestorNode, found := treePos[ancestorNum.String()]
		if !found {
			return false
		}

		if remainder.Cmp(ONE) == 0 {
			// right node
			if node.Val <= ancestorNode.Val {
				return false
			}
		} else {
			if node.Val >= ancestorNode.Val {
				return false
			}
		}

		ancestorNum, remainder = ancestorNum.DivMod(ancestorNum, TWO, remainder)
	}
	return true
}
