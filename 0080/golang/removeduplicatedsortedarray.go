package removeduplicatedsortedarray

// import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	currentNum := nums[0]
	numCount := 1 // count occurence of current number
	currIdx := 1  // current number should be put to which position

	for _, num := range nums[1:len(nums)] {
		if num == currentNum {
			if numCount >= 2 {
				continue
			}
			numCount++
		} else {
			currentNum = num
			numCount = 1
		}

		nums[currIdx] = num
		currIdx++
	}
	return currIdx
}
