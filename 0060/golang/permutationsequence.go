package permutationsequence

import "fmt"

func getPermutation(n int, k int) string {
	targetIndex := k - 1 // from 1-based to 0-bzsed

	if n == 1 && k == 1 {
		return "1"
	}

	// init slice
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	subNums := nums
	var numIdx int
	remain := targetIndex
	sizeOfSubNumsPermutation := factioral(n - 1)
	permutations := []int{}
	for len(subNums) > 1 {
		fmt.Printf("subNums: %+v\n", subNums)
		numIdx = remain / sizeOfSubNumsPermutation
		fmt.Printf("numIdx: %d, remain: %d, sizeOfSubNumsPermutation: %d\n", numIdx, remain, sizeOfSubNumsPermutation)
		remain = remain % sizeOfSubNumsPermutation
		permutations = append(permutations, subNums[numIdx])

		subNums = getSubNums(numIdx, subNums)
		sizeOfSubNumsPermutation = factioral(len(subNums) - 1)
	}

	permutations = append(permutations, subNums[0])

	target := ""
	for _, num := range permutations {
		target = fmt.Sprintf("%s%d", target, num)
	}

	return target
}

func getSubNums(idx int, nums []int) []int {
	var subNums []int
	if idx == 0 {
		subNums = nums[1:len(nums)]
	} else if idx == len(nums)-1 {
		subNums = nums[0:idx]
	} else {
		subNums = []int{}
		subNums = append(subNums, nums[0:idx]...)
		subNums = append(subNums, nums[idx+1:len(nums)]...)
	}
	return subNums
}

func factioral(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	num := 1
	for i := 2; i <= n; i++ {
		num *= i
	}
	return num
}
