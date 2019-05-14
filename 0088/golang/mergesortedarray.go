package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	// for each array, rank then from large to small
	nums1Ind := m - 1
	nums2Ind := n - 1
	ind := m + n - 1

	for ind >= 0 {
		if nums1Ind < 0 {
			nums1[ind] = nums2[nums2Ind]
			nums2Ind--
		} else if nums2Ind < 0 {
			nums1[ind] = nums1[nums1Ind]
			nums1Ind--
		} else if nums1[nums1Ind] <= nums2[nums2Ind] {
			nums1[ind] = nums2[nums2Ind]
			nums2Ind--
		} else {
			nums1[ind] = nums1[nums1Ind]
			nums1Ind--
		}

		ind--
	}
}
