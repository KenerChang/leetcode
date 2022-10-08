package relativesortarray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// Can solve this problem with concept of counting sort
	// we can use a map to keep count of each element in arr1
	// and then iterate over arr2 to put elements with corresponding count
	// the put elements not in arr2 with ancending order

	arr2Nums := map[int]bool{}
	for _, num := range arr2 {
		arr2Nums[num] = true
	}

	counts := map[int]int{}
	notInArra2 := []int{}
	for _, v := range arr1 {
		counts[v]++

		if _, ok := arr2Nums[v]; !ok {
			notInArra2 = append(notInArra2, v)
		}
	}

	sort.Ints(notInArra2)

	results := []int{}
	for _, num := range arr2 {
		for i := 0; i < counts[num]; i++ {
			results = append(results, num)
		}
	}

	results = append(results, notInArra2...)

	return results
}
