package maximumlengthofrepeatedsubarray

func findLength(nums1 []int, nums2 []int) int {
	// we can use dp to solve this problem
	// transition formula:
	// dp[i][j] = dp[i-1][j-1] +1 // if nums1[i] == nums2[j]
	// dp[i][j] = 0 // consecutive stop

	// init dp
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}

	var result int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if i == 0 || j == 0 {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1
				}
			} else {
				if nums1[i] == nums2[j] {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			result = max(result, dp[i][j])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
