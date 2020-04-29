package uniquebinarysearchtreesii

import (
	btit "leetcode/0094/golang/binarytreeinordertraversal"
	"strconv"
)

func generateTrees(n int) []*btit.TreeNode {
	if n == 0 {
		return []*btit.TreeNode{}
	}

	cache := map[string][]*btit.TreeNode{}
	return generateTreesImpl(1, n, cache)
}

func generateTreesImpl(start, end int, cache map[string][]*btit.TreeNode) (trees []*btit.TreeNode) {
	if start == end {
		trees = []*btit.TreeNode{
			&btit.TreeNode{
				Val: start,
			},
		}
		return
	}

	key := toKey(start, end)
	if trees, found := cache[key]; found {
		return trees
	}

	for num := start; num <= end; num++ {
		if num == start {
			// since nums are sorted, first number can only has right node
			subTrees := generateTreesImpl(start+1, end, cache)
			for _, subTree := range subTrees {
				trees = append(trees, &btit.TreeNode{
					Val:   num,
					Right: subTree,
				})
			}
		} else if num == end {
			// last number can only has left node
			subTrees := generateTreesImpl(start, end-1, cache)
			for _, subTree := range subTrees {
				trees = append(trees, &btit.TreeNode{
					Val:  num,
					Left: subTree,
				})
			}
		} else {
			leftSubTrees := generateTreesImpl(start, num-1, cache)
			rightSubTrees := generateTreesImpl(num+1, end, cache)
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

func toKey(start, end int) string {
	return strconv.Itoa(start) + "," + strconv.Itoa(end)
}
