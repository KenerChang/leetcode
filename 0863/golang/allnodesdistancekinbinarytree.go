package allnodesdistancekinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func collectNodesInK(root *TreeNode, k int) []int {
	if root == nil {
		return []int{}
	}

	// fmt.Printf("root.Val: %d, k: %d\n", root.Val, k)
	if k == 0 {
		return []int{root.Val}
	}

	results := collectNodesInK(root.Left, k-1)
	results = append(results, collectNodesInK(root.Right, k-1)...)

	return results
}

func dfsUtil(root *TreeNode, target *TreeNode, k int) ([]int, int) {
	if root == nil {
		return []int{}, -1
	}

	var results []int
	if root == target {
		if k == 0 {
			return []int{root.Val}, 0
		}

		// find child nodes in k distance
		results = collectNodesInK(root.Left, k-1)
		results = append(results, collectNodesInK(root.Right, k-1)...)
		return results, 0
	}

	results, childDistanceFromTarget := dfsUtil(root.Left, target, k)
	distanceFromTarget := childDistanceFromTarget + 1
	if childDistanceFromTarget != -1 {
		// fmt.Printf("left root.Val: %d, distanceFromTarget: %d\n", root.Val, distanceFromTarget)
		if distanceFromTarget == k {
			results = append(results, root.Val)
		} else if distanceFromTarget < k {
			results = append(results, collectNodesInK(root.Right, k-distanceFromTarget-1)...)
		}
		return results, distanceFromTarget
	}

	results, childDistanceFromTarget = dfsUtil(root.Right, target, k)
	distanceFromTarget = childDistanceFromTarget + 1
	if childDistanceFromTarget != -1 {
		// fmt.Printf("right root.Val: %d, distanceFromTarget: %d\n", root.Val, distanceFromTarget)
		if distanceFromTarget == k {
			results = append(results, root.Val)
		} else if distanceFromTarget < k {
			results = append(results, collectNodesInK(root.Left, k-distanceFromTarget-1)...)
		}
		return results, distanceFromTarget
	}

	return []int{}, -1
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// Can use DFS to find the target, once the target is found,
	// then we start to find nodes in k distance

	results, _ := dfsUtil(root, target, k)

	return results
}
