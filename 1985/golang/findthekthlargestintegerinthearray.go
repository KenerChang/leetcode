package findthekthlargestintegerinthearray

func gte(num1, num2 string) bool {
	if len(num1) != len(num2) {
		return len(num1) > len(num2)
	}

	return num1 > num2
}

func quickSelect(nums []string, l, r, k int) {
	if l >= r {
		return
	}
	pivot := nums[r]

	idx := l
	for i := l; i < r; i++ {
		if gte(nums[i], pivot) {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}

	nums[idx], nums[r] = nums[r], nums[idx]
	// fmt.Printf("nums: %+v, l: %d, r: %d, k: %d, idx: %d, count: %d\n", nums, l, r, k, idx, idx-l+1)

	count := idx - l + 1
	if count == k {
		return
	} else if count < k {
		quickSelect(nums, idx+1, r, k-count)
	} else {
		quickSelect(nums, l, idx-1, k)
	}
}

func kthLargestNumber(nums []string, k int) string {
	quickSelect(nums, 0, len(nums)-1, k)

	// fmt.Printf("nums: %+v, k: %d\n", nums, k)

	min := nums[0]
	for i := 0; i < k; i++ {
		if gte(min, nums[i]) {
			min = nums[i]
		}
	}
	return min
}
