package querykthsmallesttrimmednumber

type node struct {
	val string
	idx int
}

func countingSort(nums []node, trimIdx int) []node {
	// use counting sort as radix sort subroutine

	counts := make([]int, 10)

	// count
	r := len(nums[0].val)
	for _, num := range nums {
		idx := r - trimIdx
		digit := num.val[idx] - '0'
		counts[digit]++
	}

	// accumulate
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// output
	outputs := make([]node, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		idx := r - trimIdx
		digit := nums[i].val[idx] - '0'
		outputs[counts[digit]-1] = nums[i]
		counts[digit]--
	}

	return outputs
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	// We can use radix sort to solve this problem.
	// During radix sort, we can keep
	// and we can answer query in O(1) time.

	// build nodes
	ns := make([]node, len(nums))
	for i, n := range nums {
		ns[i] = node{n, i}
	}

	// radix sort
	r := len(nums[0])
	mins := map[int][]node{}
	for i := r - 1; i >= 0; i-- {
		trimIdx := r - i

		// counting sort
		ns = countingSort(ns, trimIdx)
		mins[trimIdx] = ns
	}

	// answer queries
	results := make([]int, len(queries))
	for i, q := range queries {
		k, trimIdx := q[0], q[1]
		results[i] = mins[trimIdx][k-1].idx
	}

	return results
}
