package containsduplicate

// import (
// 	"sort"
// )

func containsDuplicate(nums []int) bool {
	cache := map[int]bool{}
	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return true
		}
		cache[num] = true
	}

	return false
}
