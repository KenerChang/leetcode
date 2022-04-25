package wigglesort

func wiggleSort(nums []int) {
	// when we process i, it means all nums before nums[i] are wiggle sorte
	// if i % 2 == 0, and nums[i] < nums[i+1], it meets the condition after swap nums[i] and nums[i+1]
	// else  if i % 2 = 1 and nums[i] > nums[i+1], it also meets the condition swap nums[i] and nums[i+1]

	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		} else {
			if nums[i+1] > nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
