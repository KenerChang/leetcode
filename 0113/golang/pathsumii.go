package pathsumii

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	resultsStr := pathSumImpl(root, sum, []string{}, "")

	results := make([][]int, len(resultsStr))
	for i, resultStr := range resultsStr {
		nums := strings.Split(resultStr, ",")

		result := make([]int, len(nums))
		for j, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			result[j] = num
		}
		results[i] = result
	}
	return results
}

func pathSumImpl(root *TreeNode, sum int, results []string, path string) []string {
	if root == nil {
		return results
	}

	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path += "," + strconv.Itoa(root.Val)
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {

		results = append(results, path)
		return results
	}

	remain := sum - root.Val
	results = pathSumImpl(root.Left, remain, results, path)
	results = pathSumImpl(root.Right, remain, results, path)
	return results
}
