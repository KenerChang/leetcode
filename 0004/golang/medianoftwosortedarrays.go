package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	middle := (total - 1) / 2
	if total%2 == 0 {
		return float64(findKth(nums1, nums2, middle)+findKth(nums1, nums2, middle+1)) / 2
	} else {
		return float64(findKth(nums1, nums2, middle))
	}
}

func findKth(nums1, nums2 []int, k int) int {
	// fmt.Printf("nums1: %+v, nums2: %+v, k: %d\n", nums1, nums2, k)
	var smaller, larger []int

	// assign smaller slice, larger slice
	if len(nums1) <= len(nums2) {
		smaller, larger = nums1, nums2
	} else {
		smaller, larger = nums2, nums1
	}

	// if size of smaller is 0, return larger[k-1]
	if len(smaller) == 0 {
		return larger[k]
	}

	// if k == 0, return min(nums1[0], nums2[0])
	if k == 0 {
		return min(smaller[0], larger[0])
	}

	// get middle index of smaller & larger
	smallerMiddle := len(smaller) / 2
	largerMiddle := len(larger) / 2

	// if k <= smallerMiddle + largerMiddle
	if k <= smallerMiddle+largerMiddle {
		// if smaller[Middle] <= largerMiddle, ignore larger right part
		if smaller[smallerMiddle] <= larger[largerMiddle] {
			larger = larger[0:largerMiddle]
		} else {
			// else ignore smaller right part
			smaller = smaller[0:smallerMiddle]
		}

	} else {
		// if smaller[middle] <= larger[largerMiddle], ignore smaller left part
		if smaller[smallerMiddle] <= larger[largerMiddle] {
			smaller = smaller[smallerMiddle+1:]
			k -= smallerMiddle + 1
		} else {
			// else ignore larger right part
			larger = larger[largerMiddle+1:]
			k -= largerMiddle + 1
		}
	}
	return findKth(smaller, larger, k)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
