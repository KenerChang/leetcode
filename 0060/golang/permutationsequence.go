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

	n_1factorial := factioral(n - 1)
	numIdx := targetIndex / n_1factorial
	shift := targetIndex % n_1factorial

	subNums := getSubNums(numIdx, nums)
	permutations := genPermutation(subNums)

	target := fmt.Sprintf("%d%s", nums[numIdx], permutations[shift])

	return target
}

func genPermutation(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{
			fmt.Sprint(nums[0]),
		}
	}

	output := []string{}
	var subNums []int
	var permutations []string
	for idx, num := range nums {
		subNums = getSubNums(idx, nums)
		permutations = genPermutation(subNums)

		for _, permutation := range permutations {
			output = append(output, fmt.Sprintf("%d%s", num, permutation))
		}
	}
	return output
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
