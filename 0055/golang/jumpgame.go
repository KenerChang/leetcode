package jumpgame

// import "fmt"

func canJump(nums []int) bool {
	// I have tried 4 implementation of jump game
	// to prevent time out of limit,
	// the only way is we reverse the path and find the first path from end to start
	// this solution can prevent time out limit of large input,
	// but it can not pass if there are many duplicated computation in the input
	// so add a cache to keep which position was computed
	// the reverse with cache version is the final implementation which can pass
	cache := map[int]bool{}
	return canJumpWithCache(nums, cache)
}

func canJumpWithCache(nums []int, cache map[int]bool) bool {
	if len(nums) == 0 {
		return false
	}

	step := nums[0]
	if step >= len(nums) || len(nums) == 1 {
		return true
	}

	result := false
	lastIndex := len(nums) - 1
	for i := lastIndex - 1; i >= 0; i-- {
		step := nums[i]
		if step+i >= lastIndex {
			if _, ok := cache[i]; !ok {
				lresult := canJumpWithCache(nums[0:i+1], cache)
				result = result || lresult
				cache[i] = result
				if result {
					break
				}
			}
		}
	}
	return result
}
