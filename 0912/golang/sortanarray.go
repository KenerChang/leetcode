package sortanarray

func sortArray(nums []int) []int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) []int {
	if l >= r {
		return []int{nums[l]}
	}

	mid := (l + r) / 2
	left := mergeSort(nums, l, mid)
	right := mergeSort(nums, mid+1, r)

	merged := merge(left, right)
	return merged
}

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	merged := []int{}

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}
