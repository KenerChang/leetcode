package wigglesortii

func findKthSmallest(nums []int, k int) int {
	heap := make([]int, k)
	copy(heap, nums)
	for i := (k - 2) / 2; i >= 0; i-- {
		helper(heap, i, k)
	}
	for i := k; i < len(nums); i++ {
		if heap[0] > nums[i] {
			heap[0] = nums[i]
			helper(heap, 0, k)
		}
	}
	return heap[0]
}

func helper(heap []int, parent, n int) {
	for child := 2*parent + 1; child < n; parent, child = child, 2*child+1 {
		if child+1 < n && heap[child] < heap[child+1] {
			child++
		}
		if heap[parent] > heap[child] {
			return
		}
		heap[parent], heap[child] = heap[child], heap[parent]
	}
}

func wiggleSort(nums []int) {
	// We can find the median of the nums.

	m := (len(nums) + 1) / 2
	median := findKthSmallest(nums, m)

	i := 0 // from head
	j := 0
	n := len(nums)
	k := len(nums) - 1 // from tail
	for j <= k {
		if nums[newIndex(j, n)] > median {
			nums[newIndex(i, n)], nums[newIndex(j, n)] = nums[newIndex(j, n)], nums[newIndex(i, n)]
			i++
			j++
		} else if nums[newIndex(j, n)] < median {
			nums[newIndex(j, n)], nums[newIndex(k, n)] = nums[newIndex(k, n)], nums[newIndex(j, n)]
			k--
		} else {
			j++
		}
	}
}

func newIndex(i, n int) int {
	return (2*i + 1) % (n | 1)
}
