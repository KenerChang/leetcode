package uniquebinarysearchtreesii

import (
	btit "leetcode/0094/golang/binarytreeinordertraversal"
	"strconv"
)

func generateTrees(n int) []*btit.TreeNode {
	if n == 0 {
		return []*btit.TreeNode{}
	}
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	cache := map[string][]*btit.TreeNode{}
	return generateTreesImpl(nums, cache)
}

func generateTreesImpl(nums []int, cache map[string][]*btit.TreeNode) (trees []*btit.TreeNode) {
	key := toKey(nums)
	if trees, found := cache[key]; found {
		return trees
	}

	if len(nums) == 1 {
		trees = []*btit.TreeNode{
			&btit.TreeNode{
				Val: nums[0],
			},
		}
		cache[key] = trees
		return
	}

	numLen := len(nums)
	for idx, num := range nums {
		if idx == 0 {
			// since nums are sorted, first number can only has right node
			subTrees := generateTreesImpl(nums[1:numLen], cache)
			for _, subTree := range subTrees {
				trees = append(trees, &btit.TreeNode{
					Val:   num,
					Right: subTree,
				})
			}
		} else if idx == numLen-1 {
			// last number can only has left node
			subTrees := generateTreesImpl(nums[0:numLen-1], cache)
			for _, subTree := range subTrees {
				trees = append(trees, &btit.TreeNode{
					Val:  num,
					Left: subTree,
				})
			}
		} else {
			leftSubTrees := generateTreesImpl(nums[0:idx], cache)
			rightSubTrees := generateTreesImpl(nums[idx+1:numLen], cache)
			for _, leftTree := range leftSubTrees {
				for _, rightTree := range rightSubTrees {
					trees = append(trees, &btit.TreeNode{
						Val:   num,
						Left:  leftTree,
						Right: rightTree,
					})
				}
			}
		}
	}
	cache[key] = trees
	return
}

func toKey(nums []int) string {
	key := ""
	for _, num := range nums {
		key += "," + strconv.Itoa(num)
	}
	return key
}
