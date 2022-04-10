package closestbinarysearchtreevalueii

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	// find closetValue
	closetValue := closestValue(root, target)
	results := []int{closetValue}
	if k == 1 {
		return results
	}

	// do inorder traveral and get int slice
	sortedVals := preorder(root, []int{})
	if k == len(sortedVals) {
		return sortedVals
	}

	// find closeValue position in this slice
	var closetValuePos int
	for i, val := range sortedVals {
		if val == closetValue {
			closetValuePos = i
			break
		}
	}

	// use two pointer to get k closet valuesnt
	ptrL, ptrR := closetValuePos-1, closetValuePos+1

	// while ptrR < N and ptrL >= 0 and len(results) < k
	for ptrL >= 0 && ptrR < len(sortedVals) && len(results) < k {
		// diffR := sortedVlas[ptrR] - target
		// diff

		diffR := math.Abs(float64(sortedVals[ptrR]) - target)
		diffL := math.Abs(float64(sortedVals[ptrL]) - target)

		// if diffR < diffL, put slice[ptrR] to results, ptrR++
		if diffR < diffL {
			results = append(results, sortedVals[ptrR])
			ptrR++
		} else {
			// else put slice[ptrL] to results, ptrL--
			results = append(results, sortedVals[ptrL])
			ptrL--
		}
	}

	// if len(results) >= k return resutls
	if len(results) >= k {
		return results
	}

	// if ptrL > 0, iterate left part until len(results) == k
	if ptrL >= 0 {
		for len(results) < k {
			results = append(results, sortedVals[ptrL])
			ptrL--
		}
	} else if ptrR < len(sortedVals) {
		// else if ptrR < N - 1, iterate right part until len(results) == k
		for len(results) < k {
			results = append(results, sortedVals[ptrR])
			ptrR++
		}
	}

	return results
}

func closestValue(root *TreeNode, target float64) int {
	node := root
	minVal := math.MaxFloat64
	var nodeVal int
	// start from root
	for node != nil {
		diff := float64(node.Val) - target
		diff = math.Abs(diff)

		if diff < minVal {
			// if node.Val - target < minVal, update minVal & nodeVal
			minVal = diff
			nodeVal = node.Val
		} else if minVal == 0 {
			// else if node.Val - target > minVal, no need to go down since the difference will become larger
			break
		}

		// if target is larger than root go right, else go left
		if float64(node.Val) < target {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return nodeVal
}

func preorder(root *TreeNode, traversal []int) []int {
	if root == nil {
		return traversal
	}

	// travel left tree
	traversal = preorder(root.Left, traversal)

	// put node val
	traversal = append(traversal, root.Val)

	// travel right tree
	traversal = preorder(root.Right, traversal)

	return traversal
}
