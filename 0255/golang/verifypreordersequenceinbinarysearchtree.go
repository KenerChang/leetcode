package verifypreordersequenceinbinarysearchtree

import "math"

func verifyPreorder(preorder []int) bool {
	// if there is only one val in preorder, it is a binary search tree sequence
	if len(preorder) == 1 {
		return true
	}

	stack := []int{preorder[0]}
	lo := math.MinInt32
	for i := 1; i < len(preorder); i++ {
		// if preorder[i] is less than lo, return false
		if preorder[i] < lo {
			return false
		}

		if preorder[i] > preorder[i-1] {
			// find the first value in stack which is larger the preorder[i], or until to root, then update lo
			for len(stack) >= 1 && stack[len(stack)-1] < preorder[i] {
				lo = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}

		stack = append(stack, preorder[i])
	}

	return true
}
