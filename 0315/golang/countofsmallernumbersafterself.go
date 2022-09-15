package countofsmallernumbersafterself

type treeNode struct {
	sorted []int
	left   *treeNode
	right  *treeNode
	l      int
	r      int
}

func build(nums []int, l, r int) *treeNode {
	if l == r {
		return &treeNode{
			sorted: []int{nums[l]},
			l:      l,
			r:      r,
		}
	}

	node := &treeNode{
		l: l,
		r: r,
	}

	mid := (l + r) / 2
	node.left = build(nums, l, mid)
	node.right = build(nums, mid+1, r)

	node.sorted = merge(node.left.sorted, node.right.sorted)
	return node
}

func merge(a, b []int) []int {
	// merge two sorted array
	var pa, pb int
	results := []int{}
	for pa < len(a) && pb < len(b) {
		if a[pa] < b[pb] {
			results = append(results, a[pa])
			pa++
		} else {
			results = append(results, b[pb])
			pb++
		}
	}

	if pa < len(a) {
		results = append(results, a[pa:]...)
	} else {
		results = append(results, b[pb:]...)
	}

	return results
}

func count(nums []int, val int) int {
	// find first position p that nums[p] >= val
	if len(nums) == 1 {
		if nums[0] >= val {
			return 0
		} else {
			return 1
		}
	}

	l, r := 0, len(nums)-1
	for l < r {
		// fmt.Printf("l=%d, r=%d, val: %d\n", l, r, val)
		mid := (l + r) / 2

		if nums[mid] >= val {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// fmt.Printf("nums: %v, val: %d, l: %d, r: %d\n", nums, val, l, r)

	if nums[l] < val {
		return l + 1
	}
	return l
}

func find(root *treeNode, l, val int) int {
	if root == nil {
		return 0
	}

	// fmt.Printf("root.l: %d, root.r: %d, l: %d, val: %d\n", root.l, root.r, l, val)

	if root.l >= l {
		return count(root.sorted, val)
	} else {
		// root.l < l, search right
		mid := (root.l + root.r) / 2
		if mid >= l {
			// search both left & right
			return find(root.left, l, val) + find(root.right, l, val)
		} else {
			return find(root.right, l, val)
		}
	}
}

func countSmaller(nums []int) []int {
	// we can use segment tree to solve this problem
	// in each node we store sorted array of nums[l..r]
	// after the segment tree is built, search nums which are larger than x of range [l..r]
	// is O(logn)
	// and for all n with index i in nums, we can find numbers which are larger than i in [i+1..size(n)]
	// we can solve it in O(nlogn)

	root := build(nums, 0, len(nums)-1)
	results := make([]int, len(nums))
	for i, num := range nums {
		results[i] = find(root, i+1, num)
	}

	return results
}
